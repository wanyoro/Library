package main

import (
	"encoding/json"
	// "errors"
	// "io/ioutil"
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
	w.Header().Set("Content-Type", "Application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		ERROR(w, http.StatusBadRequest, err)
		return
	}
	userGotten, err := GetUserWithId(id, a.DB)
	if err != nil {
		ERROR(w, http.StatusBadRequest, err)
		return
	}
	JSON(w, http.StatusOK, userGotten)
}

// func (a *App) UpdateUser(w http.ResponseWriter, r *http.Request) {

// 	vars := mux.Vars(r)
// 	uid, err := strconv.ParseUint(vars["id"], 10, 32)
// 	if err != nil {
// 		ERROR(w, http.StatusBadRequest, err)
// 		return
// 	}
// 	body, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		ERROR(w, http.StatusUnprocessableEntity, err)
// 		return
// 	}
// 	user := Person{}
// 	err = json.Unmarshal(body, &user)
// 	if err != nil {
// 		ERROR(w, http.StatusUnprocessableEntity, err)
// 	}
// 	tokenID, err := ExtractTokenID(r)
// 	if err != nil {
// 		ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
// 		return
// 	}
// 	if tokenID != uint32(uid) {
// 		ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
// 		return
// 	}
// 	updatedUser, err := user.UpdateUser(a.DB, uint32(uid))
// 	if err!= nil{
// 		formattedError := FormatError(err.Error())
// 		ERROR(w, http.StatusInternalServerError, formattedError)
// 	}
// 	JSON(w, http.StatusOK, updatedUser)

// }
