package main

import (
	"encoding/json"
	"log"
	"net/http"
	"math/rand"
	"strconv"
	"github.com/gorilla/mux"
)

//Models
//Book Strcut
type Book struct {
	ID 			string  `json:"id"`
	Isbn 		string  `json:"isbn"`
	Title 	string  `json:"title"`
	Author 	*Author  `json:"author"`

}

//Author Strcut
type Author struct{
	Firstname	string `json:"firstname"`
	Lastname	string `json:"lastname"`
}


// Init Books var as a slice Book struct
var books []Book


//Get all books
func getBooks(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "Application/json")
	json.NewEncoder(w).Encode(books)
}

//Get one book
func getBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "Application/json")
	params:= mux.Vars(r) //Get Params 

	//Loop through books and find ID

	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return 
		}
	}

	json.NewEncoder(w).Encode(&Book{})
}

//createBook book
func createBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "Application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(1000000)) // Mock Id not safe for production
	books = append(books, book)

	json.NewEncoder(w).Encode(book)
}

//Update book
func updateBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "Application/json")
	params := mux.Vars(r)

	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var newBook Book
			_ = json.NewDecoder(r.Body).Decode(&newBook)
			newBook.ID = params["id"]
			books = append(books, newBook)
			json.NewEncoder(w).Encode(newBook)
			return 
		}
	}
}

//Delete book
func deleteBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)

	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}

	json.NewEncoder(w).Encode(books)
}


func main(){

	// Init route
	r := mux.NewRouter()
	
	//Mock Data -@todo - Implement Data base
	books = append(books, Book{ID: "1", Isbn: "448743", Title: "Book One", Author: &Author{Firstname: "John", Lastname: "Doe"}})
	books = append(books, Book{ID: "2", Isbn: "542865", Title: "Book Two", Author: &Author{Firstname: "Luis", Lastname: "Sandoval"}})

	//Route handlers / End points
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}