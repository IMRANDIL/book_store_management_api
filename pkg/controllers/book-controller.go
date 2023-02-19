package controllers

import (
	"book_store_api/pkg/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)



var NewBook models.Book


func GetAllBook(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBook()

	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookId := params["bookId"]

	ID, err := strconv.ParseInt(bookId,0,0)

	if err != nil {
		fmt.Println("error while parsing")
	}

	bookDetails, _ := models.GetBookById(ID)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}