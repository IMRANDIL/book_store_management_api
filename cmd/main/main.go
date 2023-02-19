package main

import (
	"book_store_api/pkg/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)



func main () {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/",r)
	log.Fatal(http.ListenAndServe("localhost:8000",r))
}