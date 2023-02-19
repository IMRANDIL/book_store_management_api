package main

import (
	"book_store_api/pkg/routes"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)



func main () {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/",r)
	fmt.Println("Server started at port :8000")
	log.Fatal(http.ListenAndServe("localhost:8000",r))
	
}