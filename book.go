package main

import (
	"gorm.io/gorm"
	//"net/http"
	//"encoding/json"
	"errors"
	"strings"
	//"fmt"
	//"io/ioutil"
)

type Book struct {
	gorm.Model
	//Id     uint   `json:"id"`
	Title    string `json:"title"`
	Isbn     string `json:"isbn"`
	Author   string `json:"author"`
	PersonID int    //`gorm:"foreignKey:PersonID"`
}

func GetAllBooks(db *gorm.DB) (*[]Book, error) {
	books := []Book{}
	if err := db.Debug().Table("books").Find(&books).Error; err != nil {
		return &[]Book{}, err
	}
	return &books, nil

}

func GetUnassignedBooks(db *gorm.DB) (*[]Book, error) {
	unassignedBooks := []Book{}
	if err := db.Debug().Where("person_id is NULL").Find(&unassignedBooks).Error; err != nil {
		return &[]Book{}, err
	}
	return &unassignedBooks, nil

}

func (b *Book) CreatedBook(db *gorm.DB) (*Book, error) {
	var err error
	err = db.Debug().Create(&b).Error
	if err != nil {
		return &Book{}, err
	}
	return b, nil
}

//w.Header().Set("Content-Type", "Application/json")
// body, err := ioutil.ReadAll(r.Body)
// if err != nil {
// 	ERROR(w, http.StatusUnprocessableEntity, err)
// }

// book := Book{}
// err = json.Unmarshal(body, &book)
// if err != nil {
// 	ERROR(w, http.StatusUnprocessableEntity, err)
// 	return
// }
// user.Prepare()
// newPassword, err := HashPassword(user.Password)
// if err != nil {
// 	ERROR(w, http.StatusInternalServerError, err)
// 	return
// }
// user.Password = newPassword

// 	err = user.Validate("")
// 	if err != nil {
// 		ERROR(w, http.StatusUnprocessableEntity, err)
// 		return
// 	}
// 	userCreated, err := book.(a.DB)
// 	if err != nil {
// 		formattedError := FormatError(err.Error())

// 		ERROR(w, http.StatusInternalServerError, formattedError)
// 		return
// 	}
// 	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, userCreated.ID))
// 	JSON(w, http.StatusCreated, userCreated)
// }
// book := Book{}
// if err := db.Debug().Table("books").Create(&book).Error; err != nil {
// 	return &Book{}, err
// }
// return &book, nil

func (b *Book) ValidateBook(action string) error {
	switch strings.ToLower(action) {
	case "create":
		if b.Title == "" {
			return errors.New("required Title")
		}
		if b.Isbn == "" {
			return errors.New("required Isbn")
		}
		if b.Author == "" {
			return errors.New("required Author")
		}
		// if b.PersonID ==""{
		// 	return errors.New("required PersonID")
		// }
	}
	return nil
}

func (b *Book) Prepare() {
	b.Author = strings.TrimSpace(b.Author)
	b.Title = strings.TrimSpace(b.Title)
	b.Isbn = strings.TrimSpace(b.Isbn)
}
