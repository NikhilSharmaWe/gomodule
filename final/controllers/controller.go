package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/NikhilSharmaWe/gomodule/final/pkg/models"
)

func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, `
	<form action="/" method="GET">
        <input type="text" name="name" placeholder="name">
        <input type="text" name="author" placeholder="author">
        <input type="text" name="publication" placeholder="publication">
        <input type="submit" name="submit" value="submit">
    </form>
    <form action="/getbooks" method="POST">
        <input type="submit" name="submit" value="getbooks">
    </form>
    <form action="/book" method="GET">
        <input type="text" name="bookId" placeholder="bookId">
        <input type="submit" name="submit" value="getbookbyid">
    </form>
	`)
	name := r.FormValue("name")
	author := r.FormValue("author")
	publication := r.FormValue("publication")
	if name == "" || author == "" || publication == "" {
		w.WriteHeader(http.StatusOK)
		return
	}
	CreateBook := &models.Book{
		Name:        name,
		Author:      author,
		Publication: publication,
	}
	err := CreateBook.CreateBook()
	if err != nil {
		handleError(w, err, "error in database while creating and adding new book", 500)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Println("Created new book")
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	newBooks, err := models.GetAllBooks()
	if err != nil {
		handleError(w, err, "error in database while getting all books", 500)
		return
	}
	res, err := json.Marshal(newBooks)
	if err != nil {
		handleError(w, err, "error while marshalling books", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	bookId := r.FormValue("bookId")
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		handleError(w, err, "error while parsing book id", 500)
		return
	}
	bookDetails, _, err := models.GetBookById(ID)
	if err != nil {
		handleError(w, err, "error in database while getting book by id", 500)
		return
	}
	res, err := json.Marshal(bookDetails)
	if err != nil {
		handleError(w, err, "error while marshalling book detail", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func handleError(w http.ResponseWriter, err error, errorMessage string, statuscode int) {
	fmt.Println(errorMessage)
	w.WriteHeader(statuscode)
	w.Write([]byte(fmt.Sprintf("%s\n%s", errorMessage, err)))
}
