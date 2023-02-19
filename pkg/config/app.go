package config

import "github.com/jinzhu/gorm"



var (
	db *gorm.DB
)

func Connect () {
	d, err := gorm.Open("mysql","root:Dil@2580123@tcp(localhost:3306)/book_store?")

	if err != nil {
		panic(err)
	}

	db = d
}