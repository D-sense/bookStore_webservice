package controllers

import (
	"database/sql"
	"encoding/json"
	"github.com/d-sense/library_web_service/helper"
	"github.com/d-sense/library_web_service/models"
	"github.com/d-sense/library_web_service/repository/book"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type Controller struct {}
var books []models.Book

func (Controller) GetBooks(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		var book models.Book
		books = []models.Book{}

		bookRepo := bookRepository.BookRepository{}
        books = bookRepo.GetBooks(db, book, books)

		helper.LogFatal(json.NewEncoder(w).Encode(&books))
	}
}

func (Controller) GetBook(db *sql.DB) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request){
		var book models.Book
		params := mux.Vars(r)
		id, err := strconv.Atoi(params["id"])
        helper.LogFatal(err)

		bookRepo := bookRepository.BookRepository{}
		book = bookRepo.GetBook(db, book, id)

		helper.LogFatal(json.NewEncoder(w).Encode(book))
	}
}

func (Controller) AddBook(db *sql.DB) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request){
		var book models.Book
		var bookID int

		helper.LogFatal(json.NewDecoder(r.Body).Decode(&book))

		bookRepo := bookRepository.BookRepository{}
		bookID = bookRepo.AddBook(db, book)

		helper.LogFatal(json.NewEncoder(w).Encode(bookID))
	}
}

func (Controller) UpdateBook(db *sql.DB) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request){
		var book models.Book
		helper.LogFatal(json.NewDecoder(r.Body).Decode(&book))

		bookRepo := bookRepository.BookRepository{}
		rowsUpdated := bookRepo.UpdateBook(db, book)

		helper.LogFatal(json.NewEncoder(w).Encode(rowsUpdated))
	}
}

func (Controller) DeleteBook(db *sql.DB) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request){
		params := mux.Vars(r)
		id, err := strconv.Atoi(params["id"])
		helper.LogFatal(err)

		bookRepo := bookRepository.BookRepository{}
		rowDeleted := bookRepo.DeleteBook(db, id)

		helper.LogFatal(json.NewEncoder(w).Encode(rowDeleted))
	}
}
