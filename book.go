package main

import (
	"errors"
	"strings"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
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

	if err := db.Debug().Create(&b).Error; err != nil {
		return &Book{}, err
	}
	return b, nil
}

func GetBookById(id int, db *gorm.DB) (*Book, error) {
	book := &Book{}
	if err := db.Debug().Table("books").Where("id = ?", id).First(book).Error; err != nil {

	}
	return book, nil
}

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

	}
	return nil
}

func (b *Book) Prepare() {
	b.Author = strings.TrimSpace(b.Author)
	b.Title = strings.TrimSpace(b.Title)
	b.Isbn = strings.TrimSpace(b.Isbn)
}

func (b *Book) DeleteBook(id int, db *gorm.DB) error {
	if err := db.Debug().Table("books").Where("id=?", id).Delete(&Book{}).Error; err != nil {
		return err
	}
	return nil
}
