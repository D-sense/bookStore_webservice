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

//func getBooks(w http.ResponseWriter, r *http.Request){
//	var book models.Book
//	books = []models.Book{}
//
//	rows, err := db.Query("SELECT * FROM books")
//	helper.LogFatal(err)
//
//	defer rows.Close()
//
//	for rows.Next() {
//		err := rows.Scan(&book.ID, &book.Author, &book.Title, &book.Year)
//		helper.LogFatal(err)
//
//		books = append(books, book)
//	}
//
//	log.Println(len(books))
//	json.NewEncoder(w).Encode(&books)
//}

//func getBook(w http.ResponseWriter, r *http.Request){
//	var book models.Book
//	params := mux.Vars(r)
//
//	row := db.QueryRow("select * from books where id=$1", params["id"])
//
//	err := row.Scan(&book.ID, &book.Author, &book.Title, &book.Year)
//	helper.LogFatal(err)
//
//	json.NewEncoder(w).Encode(book)
//}

//func addBook(w http.ResponseWriter, r *http.Request){
//	var book models.Book
//	var bookId int
//
//	json.NewDecoder(r.Body).Decode(&book)
//
//	err:= db.QueryRow("insert into books (title, author, year) values($1, $2, $3) RETURNING id;", book.Title, book.Author, book.Year).Scan(&bookId)
//	helper.LogFatal(err)
//
//	json.NewEncoder(w).Encode(&bookId)
//}

//func updateBook(w http.ResponseWriter, r *http.Request){
//	var book models.Book
//	json.NewDecoder(r.Body).Decode(&book)
//
//	result, err := db.Exec("update books set title=$1, author=$2, year=$3 where id=$4 RETURNING id", &book.Title, &book.Author, &book.Year, &book.ID)
//
//	rowsUpdated, err := result.RowsAffected()
//	helper.LogFatal(err)
//
//	json.NewEncoder(w).Encode(rowsUpdated)
//}

//func deleteBook(w http.ResponseWriter, r *http.Request){
//	params := mux.Vars(r)
//
//	result, err := db.Exec("delete from books where id=$1", params["id"])
//	helper.LogFatal(err)
//
//	rowDeleted, err := result.RowsAffected()
//	helper.LogFatal(err)
//
//	json.NewEncoder(w).Encode(rowDeleted)
//}

