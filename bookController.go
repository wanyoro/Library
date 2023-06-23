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

func (a *App) CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content_Type", "Application/json")
	user := Person{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		ERROR(w, http.StatusUnprocessableEntity, err)
	}
	book := Book{}
	err = json.Unmarshal(body, &book)
	if err != nil {
		ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	book.Prepare()
	err = book.ValidateBook("")
	if err != nil {
		ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	bookCreated, err := book.CreatedBook(a.DB)
	if err != nil {
		formattedError := user.FormatError(err.Error())

		ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, bookCreated.ID))
	JSON(w, http.StatusCreated, bookCreated)

}

func (a *App) GetBookWithId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		ERROR(w, http.StatusBadRequest, err)
		return
	}
	bookGotten, err := GetBookById(id, a.DB)
	if err != nil {
		ERROR(w, http.StatusBadRequest, err)
		return
	}
	JSON(w, http.StatusOK, bookGotten)
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

func (a *App) GetUnAssignedBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	unassignedBooks, err := GetUnassignedBooks(a.DB)
	if err != nil {
		ERROR(w, http.StatusInternalServerError, err)
		return
	}
	JSON(w, http.StatusOK, unassignedBooks)
	json.NewDecoder(r.Body).Decode(&unassignedBooks)

}

func GetAllAssignedBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	var db *gorm.DB
	var assignedbks []Book
	db.Where("person_id >?", 0).Find(&assignedbks)
	json.NewEncoder(w).Encode(&assignedbks)
}

func (a *App) DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	var resp = map[string]interface{}{"status": "success", "message": "Book deleted successfully"}

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		ERROR(w, http.StatusBadRequest, err)
		return
	}

	bookGotten, err := GetBookById(id, a.DB)
	if err != nil {
		ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	err = bookGotten.DeleteBook(id, a.DB)
	if err != nil {
		ERROR(w, http.StatusInternalServerError, err)
		return
	}
	JSON(w, http.StatusOK, resp)
	return

}
