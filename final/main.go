package main

import (
	"log"
	"net/http"

	"github.com/NikhilSharmaWe/gomodule/final/controllers"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.CreateBook).Methods("GET")
	r.HandleFunc("/getbooks", controllers.GetBooks).Methods("POST")
	r.HandleFunc("/book", controllers.GetBookById).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", r))
}
