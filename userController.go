package main

import (
	"encoding/json"
	"io/ioutil"
	"strconv"

	"net/http"

	"fmt"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func (a *App) GetAllPersons(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	users, err := GetAllPersons(a.DB)
	if err != nil {
		ERROR(w, http.StatusInternalServerError, err)
	}
	JSON(w, http.StatusOK, users)
	return

	// db.First(&person, params["id"])
	// json.NewDecoder(r.Body).Decode(&person)
	// db.Save(&person)
	//json.NewEncoder(w).Encode(person)
}

func GetUserWithBkId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	var db *gorm.DB
	person_id := mux.Vars(r)
	var person Book
	err := db.Debug().Preload("Person").Where("person_id = ?", person_id["person_id"]).Find(&person).Error
	if err != nil {
		fmt.Println(err)
	}
	json.NewEncoder(w).Encode(person)
}

func GetPeopleAndBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	var db *gorm.DB
	var psn []Person
	err := db.Debug().Preload("Book").
		Joins("INNER JOIN books ON books.person_id = people.id").Find(&psn).Error
	if err != nil {
		fmt.Println(err)
	}
	json.NewEncoder(w).Encode(&psn)

}

func (a *App) GetAllUsersWithoutBks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	users, err := GetAllUsersWithoutBks(a.DB)
	if err != nil {
		ERROR(w, http.StatusInternalServerError, err)
	}
	JSON(w, http.StatusOK, users)
	return
}

func (a *App) GetUserWithId(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	uid, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		ERROR(w, http.StatusBadRequest, err)
		return
	}
	user := Person{}
	userGotten, err := user.GetUserWithId(a.DB, uint32(uid))
	if err != nil {
		ERROR(w, http.StatusBadRequest, err)
		return
	}
	JSON(w, http.StatusOK, userGotten)
}

func (a *App) CreateUser(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		ERROR(w, http.StatusUnprocessableEntity, err)
	}

	user := Person{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	user.Prepare()
	err = user.Validate("")
	if err != nil {
		ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	userCreated, err := user.SavePerson(a.DB)
	if err != nil {
		formattedError := FormatError(err.Error())

		ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, userCreated.ID))
	JSON(w, http.StatusCreated, userCreated)
}
