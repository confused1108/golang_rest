package main

import(
	"encoding/json"
	"log"
	"net/http"
	"math/rand"
	"strconv"
	"github.com/gorilla/mux"
)

//Book Structure
type Book struct{
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Author *Author `json:"author"`
}

//Author structure
type Author struct {
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

//Init books var as a slice book struct
var books []Book

//Get All Books
func getBooks(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(books)
}

//Get single book
func getBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params:=mux.Vars(r)

	//looping
	for _, item:=range books {
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})

}

//Create a record
func createBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	var book Book
	_=json.NewDecoder(r.Body).Decode(&book)
	book.ID=strconv.Itoa(rand.Intn(1000000))   //Mock
	books =append(books,book)
	json.NewEncoder(w).Encode(&Book{})
}

//update
func updateBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params:=mux.Vars(r)
	for index, item:=range books{
		if item.ID==params["id"]{
			books=append(books[:index],books[index+1:]...)
			var book Book
			_=json.NewDecoder(r.Body).Decode(&book)
			book.ID= params["id"]  //Mock
			books =append(books,book)
			json.NewEncoder(w).Encode(&Book{})
			return
		}
	} 
	json.NewEncoder(w).Encode(books)
}

//delete
func deleteBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params:=mux.Vars(r)
	for index, item:=range books{
		if item.ID==params["id"]{
			books=append(books[:index],books[index+1:]...)
			break
		}
	} 
	json.NewEncoder(w).Encode(books)
}


func main(){
	//Init Router
	r:=mux.NewRouter()

	//Mock Data
	books=append(books, Book{ID:"1",Isbn:"4852794",Title:"Book one", Author:&Author {Firstname:"John",Lastname:"Doe"}})
	books=append(books, Book{ID:"2",Isbn:"4854574",Title:"Book two", Author:&Author {Firstname:"Steve",Lastname:"Gates"}})

	//endpoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8004",r))
}
