package main

import (
	"encoding/json"
	"net/http"

	"gorm.io/gorm"
)

func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content_Type", "Application/json")
	var db *gorm.DB
	var Newbook Book
	json.NewDecoder(r.Body).Decode(&Newbook)
	db.Create(&Newbook)
	json.NewEncoder(w).Encode(&Newbook)
}

func (a *App) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	books, err := GetAllBooks(a.DB)
	if err != nil {
		ERROR(w, http.StatusInternalServerError, err)
		return
	}
	JSON(w, http.StatusOK, books)
	json.NewDecoder(r.Body).Decode(&books)

}

func GetUnAssignedBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	var db *gorm.DB
	var users []Book
	db.Where("person_id is NULL").Find(&users)
	json.NewEncoder(w).Encode(users)
}

func GetAllAssignedBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	var db *gorm.DB
	var assignedbks []Book
	db.Where("person_id >?", 0).Find(&assignedbks)
	json.NewEncoder(w).Encode(&assignedbks)
}
