package main

import (
	"encoding/json"
	//"errors"
	//"io/ioutil"
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

func (a *App) GetPeopleAndBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	pnb, err := GetPeopleAndBooks(a.DB)
	if err != nil {
		ERROR(w, http.StatusInternalServerError, err)
	}
	JSON(w, http.StatusOK, pnb)
	return
	// var db *gorm.DB
	// var psn []Person
	// err :=
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// json.NewEncoder(w).Encode(&psn)

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
// 	var resp = map[string]interface{}{"status": "success", "message": "user updated successfully"}
// 	w.Header().Set("Content-Type", "Application/json")
// 	vars := mux.Vars(r)

// 	//user := r.Context().Value("personID").(float64)
// 	//userID := uint(user)

// 	id, _ := strconv.Atoi(vars["id"])

// 	updateUser, err := GetUserWithId(id, a.DB)

// 	// if updateUser.ID != userID {
// 	// 	resp["status"] = "failed"
// 	// 	resp["message"] = "Unautorized user update"
// 	// 	JSON(w, http.StatusUnauthorized, resp)
// 	// 	return
// 	// }

// 	body, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		ERROR(w, http.StatusBadRequest, err)
// 		return
// 	}

// 	updatedUser := Person{}
// 	if err = json.Unmarshal(body, &updatedUser); err != nil {
// 		ERROR(w, http.StatusBadRequest, err)
// 		return
// 	}

// 	updatedUser.Prepare()

// 	_, err = updatedUser.UpdateUser(id, a.DB)
// 	if err != nil {
// 		ERROR(w, http.StatusInternalServerError, err)
// 		return
// 	}

// 	JSON(w, http.StatusOK, resp)
// 	return

// }
