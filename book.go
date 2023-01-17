package main

import (
	"gorm.io/gorm"
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
