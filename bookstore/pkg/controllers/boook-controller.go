package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/souravagrawal29/go-playground/bookstore/pkg/models"
	"github.com/souravagrawal29/go-playground/bookstore/pkg/utils"
)



func CreateBook(w http.ResponseWriter, r *http.Request) {
	Book := &models.Book{}
	utils.ParseBody(r, Book)
	Book, err := Book.CreateBook()
	if err != nil {
		log.Println(err)
		utils.SetErrorResponse(w, err)
	}
	utils.SetSuccessResponse(w, Book)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	bookId := mux.Vars(r)["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		log.Println(err)
		utils.SetErrorResponse(w, err)
	}
	Book := &models.Book{}
	utils.ParseBody(r, Book)
	Book, err = Book.UpdateBook(ID)
	if err != nil {
		log.Println(err)
		utils.SetErrorResponse(w, err)
	}
	utils.SetSuccessResponse(w, Book)
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	Books, err := models.GetAllBooks()
	if err != nil {
		log.Println(err)
		utils.SetErrorResponse(w, err)
		return
	}
	utils.SetSuccessResponse(w, Books)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	bookId := mux.Vars(r)["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		log.Println(err)
		utils.SetErrorResponse(w, err)
	}
	Book, err := models.GetBookById(ID)
	if err != nil {
		log.Println(err)
		utils.SetErrorResponse(w, err)
	}
	utils.SetSuccessResponse(w, Book)
}


func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0 ,0)
	if err != nil {
		log.Println(err)
		utils.SetErrorResponse(w, err)
	}
	Book, err := models.DeleteBook(ID)
	if err != nil {
		log.Println(err)
		utils.SetErrorResponse(w, err)
	}
	utils.SetSuccessResponse(w, Book)
}
