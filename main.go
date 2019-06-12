package main

import (
	"database/sql"
	"github.com/d-sense/library_web_service/controllers"
	"github.com/d-sense/library_web_service/driver"
	"github.com/d-sense/library_web_service/helper"
	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
	"net/http"
)


var db *sql.DB

func init() {
	helper.LogFatal(gotenv.Load())
}

func main(){
	db = driver.ConnectDB()

	controller := controllers.Controller{}

	router := mux.NewRouter()

	router.HandleFunc("/books", controller.GetBooks(db)).Methods("GET")
	router.HandleFunc("/book/{id}", controller.GetBook(db)).Methods("GET")
	router.HandleFunc("/book/create", controller.AddBook(db)).Methods("POST")
	router.HandleFunc("/book/update", controller.UpdateBook(db)).Methods("PUT")
	router.HandleFunc("/book/delete/{id}", controller.DeleteBook(db)).Methods("DELETE")

	println("Server started on port 80")
	helper.LogFatal(http.ListenAndServe(":80", router))
}



