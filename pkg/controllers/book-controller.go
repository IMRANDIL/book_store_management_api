package controllers

import (
	"book_store_api/pkg/models"
	"encoding/json"
	"net/http"

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
}