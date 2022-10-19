package controllers

import (
	"encoding/json"
	"fmt"
	"go_book_management_system/pkg/models"
	"go_book_management_system/pkg/utils"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

var NewBook models.Book

func CreateBook(response http.ResponseWriter, request *http.Request) {
	createBook := &models.Book{Model: gorm.Model{ID: uint(rand.Uint32())}, Name: "random", Author: "henry", Publication: "string"}
	utils.ParseBody(request, createBook)
	book := createBook.CreateBook()
	results, _ := json.Marshal(book)

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	response.Write(results)
}

func GetBook(response http.ResponseWriter, request *http.Request) {
	newBooks := models.GetAllBooks()
	results, _ := json.Marshal(newBooks)

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	response.Write(results)
}

func GetBookById(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	bookId := params["id"]

	id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	bookDetails, _ := models.GetBookById(id)

	results, _ := json.Marshal(bookDetails)

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	response.Write(results)
}

func UpdateBook(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	bookId, err := strconv.ParseInt(params["id"], 0, 0)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		response.WriteHeader(http.StatusInternalServerError)
		return
	}
	updateBook := &models.Book{}
	utils.ParseBody(request, updateBook)
	models.UpdateBook(bookId, updateBook.Name, updateBook.Author, updateBook.Publication)

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
}

func DeleteBook(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	bookId := params["id"]
	id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		response.WriteHeader(http.StatusInternalServerError)
		return
	}
	book := models.DeleteBook(id)

	results, _ := json.Marshal(book)

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	response.Write(results)
}
