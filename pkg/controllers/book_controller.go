package controllers

import (
	"bookstore_management/pkg/models"
	"bookstore_management/pkg/utils"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var NewBook models.Book

func GetBooks(w http.ResponseWriter, r *http.Request) {
	books := models.GetAllBooks()
	res, _ := json.Marshal(books)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	book_id := params["id"]
	ID, err := strconv.ParseInt(book_id, 0, 0)
	if err != nil {
		http.Error(w, "Invalid Book ID", http.StatusBadRequest)
		return
	}
	book, _ := models.GetBookById(ID)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	book := &models.Book{}
	utils.ParseBody(r, book)
	b := book.CreateBook()
	res, _ := json.Marshal(b)
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	book_id := params["id"]
	ID, err := strconv.ParseInt(book_id, 0, 0)
	if err != nil {
		http.Error(w, "Invalid Book ID", http.StatusBadRequest)
		return
	}
	book := models.DeleteBook(ID)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var book = &models.Book{}
	utils.ParseBody(r, book)
	params := mux.Vars(r)
	book_id := params["id"]
	ID, err := strconv.ParseInt(book_id, 0, 0)
	if err != nil {
		http.Error(w, "Invalid Book ID", http.StatusBadRequest)
		return
	}
	book_details, db := models.GetBookById(ID)
	if book.Name != "" {
		book_details.Name = book.Name
	}
	if book.Author != "" {
		book_details.Author = book.Author
	}
	if book.Publication != "" {
		book_details.Publication = book.Publication
	}
	db.Save(&book)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
