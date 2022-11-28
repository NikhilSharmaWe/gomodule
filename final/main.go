package main

import (
	"log"
	"net/http"

	"github.com/NikhilSharmaWe/gomodule/final/pkg/handle"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handle.CreateBook).Methods("GET")
	r.HandleFunc("/getbooks", handle.GetBooks).Methods("POST")
	r.HandleFunc("/book", handle.GetBookById).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", r))
}
