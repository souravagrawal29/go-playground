package models

import (
	"github.com/souravagrawal29/go-playground/bookstore/pkg/config"
	"gorm.io/gorm"
)


var db *gorm.DB

type Book struct {
	gorm.Model
	Name string `json:"name"`
	Author string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDb()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() (*Book, error) {
	if err := db.Create(b).Error; err != nil {
		return nil, err
	}
	return b, nil
}

func (b *Book) UpdateBook(Id int64) (*Book, error) {
	if err := db.Where("ID=?", Id).Updates(b).Error; err != nil {
		return nil, err
	}
	return b, nil
}

func GetAllBooks() ([]Book, error) {
	var Books []Book
	if err := db.Find(&Books).Error; err != nil {
		return nil, err
	}
	return Books, nil
}

func GetBookById(Id int64) (*Book, error) {
	var Book Book
	if err := db.Where("ID=?", Id).Find(&Book).Error; err != nil {
		return nil, err
	}
	return &Book, nil
}



func DeleteBook(Id int64) (*Book, error) {
	var Book Book
	if err := db.Where("ID=?", Id).Delete(&Book).Error; err != nil {
		return nil, err
	}
	return &Book, nil
}
