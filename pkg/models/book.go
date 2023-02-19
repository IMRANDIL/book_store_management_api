package models

import (
	"book_store_api/pkg/config"

	"github.com/jinzhu/gorm"
)




var db *gorm.DB

type Book struct {
	gorm.Model
	Name string `gorm:"" json:"name"`
	Author string `json:"author"`
	Publication string `json:"publication"`

}

func init () {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}


func (b *Book) CreateBook() *Book{
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBook() []Book{

	var Books []Book
	db.Find(&Books)
	return Books

}

