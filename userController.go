package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

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

func (a *App) GetAllUsersWithBks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	users, err := GetAllUsersWithBks(a.DB)
	if err != nil {
		ERROR(w, http.StatusInternalServerError, err)
	}
	JSON(w, http.StatusOK, users)
	return
}

func (a *App) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var resp = map[string]interface{}{"status": "success", "message": "User updated successfully"}
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		ERROR(w, http.StatusBadRequest, err)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	userGotten, err := GetUserWithId(id, a.DB)
	if err != nil {
		ERROR(w, http.StatusBadRequest, err)
		return
	}
	err = json.Unmarshal(body, &userGotten)
	if err != nil {
		ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	newPassword, err := HashPassword(userGotten.Password)
	if err != nil {
		ERROR(w, http.StatusInternalServerError, err)
		return
	}
	userGotten.Password = newPassword
	_, err = userGotten.UpdateUser(id, a.DB)
	if err != nil {
		ERROR(w, http.StatusInternalServerError, err)
		return
	}
	JSON(w, http.StatusOK, resp)
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
