package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//GET ALL BOOKS

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // as its a good practice to not return it as plain text
	json.NewEncoder(w).Encode(books)                   // write to response and encode books to json

}

// GET BOOK
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

// CREATE BOOK
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(10000000))
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

//UPDATE BOOK
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = item.ID
			books = append(books, book)
			json.NewEncoder(w).Encode(books)
			return
		}
	}
	json.NewEncoder(w).Encode(books)
}

//delete book
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}

// Book struct (MODEL)

type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"Isbn"`
	Title  string  `json:"Title"`
	Author *Author `json:"author"`
}
type Author struct {
	Firstname string `json:firstname`
	Lastname  string `json:lastname`
}

//Init book slice
var books []Book

func main() {
	//Init Router
	r := mux.NewRouter()

	// Mock data
	books = append(books, Book{ID: "1", Isbn: "6727672", Title: "Book one", Author: &Author{
		Firstname: "john",
		Lastname:  "doe",
	}})
	books = append(books, Book{ID: "2", Isbn: "67787672", Title: "Book two", Author: &Author{
		Firstname: "steve",
		Lastname:  "mith",
	}})
	//Route Handlers
	//End points
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/book/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")
	if err := http.ListenAndServe("localhost:8000", r); err != nil {
		log.Fatal(err)
	}

}
