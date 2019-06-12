package bookRepository

import (
	"database/sql"
	"github.com/d-sense/library_web_service/helper"
	"github.com/d-sense/library_web_service/models"
)

type BookRepository struct {}

const (
	bookStore = "book_store"
)

func (BookRepository) GetBooks(db *sql.DB, book models.Book, books []models.Book) []models.Book {
	rows, err := db.Query("SELECT * FROM " + bookStore)
	helper.LogFatal(err)

	defer helper.LogFatal(rows.Close())

	for rows.Next() {
		err := rows.Scan(&book.ID, &book.Author, &book.Title, &book.Year, &book.Category, &book.Publisher, &book.ISBN)
		helper.LogFatal(err)

		books = append(books, book)
	}

	return books
}

func (BookRepository) GetBook(db *sql.DB, book models.Book, id int) models.Book {
	row := db.QueryRow("select * from " + bookStore + " where id=$1", id)

	err := row.Scan(&book.ID, &book.Author, &book.Title, &book.Year, &book.Category, &book.Publisher, &book.ISBN)
	helper.LogFatal(err)

	return book
}

func (BookRepository) AddBook(db *sql.DB, book models.Book) int {
	err:= db.QueryRow("insert into " + bookStore + " (title, author, year, category, publisher, isbn) values($1, $2, $3, $4, $5, $6) RETURNING id;", book.Title, book.Author, book.Year, &book.ISBN, &book.Category, &book.Publisher).Scan(&book.ID)
	helper.LogFatal(err)

	return book.ID
}

func (BookRepository) UpdateBook(db *sql.DB, book models.Book) int64 {
	result, err := db.Exec("update " + bookStore + " set title=$1, author=$2, year=$3, category=$4, publisher=$5, isbn=$6 where id=$7 RETURNING id", &book.Title, &book.Author, &book.Year, &book.Category, &book.Publisher, &book.ISBN, &book.ID)

	rowsUpdated, err := result.RowsAffected()
	helper.LogFatal(err)

	return rowsUpdated
}

func (BookRepository) DeleteBook(db *sql.DB, id int) int64 {
	result, err := db.Exec("delete from " + bookStore + " where id=$1", id)
	helper.LogFatal(err)

	rowDeleted, err := result.RowsAffected()
	helper.LogFatal(err)

	return rowDeleted
}

