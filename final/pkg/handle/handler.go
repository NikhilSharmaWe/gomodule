package handle

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/NikhilSharmaWe/gomodule/final/pkg/database"
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
	CreateBook := &database.Book{
		Name:        name,
		Author:      author,
		Publication: publication,
	}
	CreateBook.CreateBook()
	w.WriteHeader(http.StatusOK)
	fmt.Println("Created new book")
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	newBooks := database.GetAllBooks()
	res, err := json.Marshal(newBooks)
	if err != nil {
		fmt.Println("error while marshalling books")
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	bookId := r.FormValue("bookId")
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing book id")
		w.WriteHeader(http.StatusInternalServerError)
	}
	bookDetails, _ := database.GetBookById(ID)
	res, err := json.Marshal(bookDetails)
	if err != nil {
		fmt.Println("error while marshalling book detail")
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
