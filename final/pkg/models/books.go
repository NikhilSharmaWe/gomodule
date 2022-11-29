package models

import (
	"github.com/NikhilSharmaWe/gomodule/final/pkg/config"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	db = config.Connect()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() error {
	db.NewRecord(b)
	db.Create(&b)
	err := db.Error
	if err != nil {
		return err
	}
	return nil
}

func GetAllBooks() ([]Book, error) {
	var Books []Book
	db.Find(&Books)
	err := db.Error
	if err != nil {
		return Books, err
	}
	return Books, nil
}

func GetBookById(Id int64) (*Book, *gorm.DB, error) {
	var getBook Book
	db := db.Where("ID=?", Id).Find(&getBook)
	err := db.Error
	if err != nil {
		return &getBook, db, err
	}
	return &getBook, db, nil
}
