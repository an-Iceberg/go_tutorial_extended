package models

import (
	"go_book_management_system/pkg/config"

	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (book *Book) CreateBook() *Book {
	db.Create(&book)
	return book
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	results := db.Where("ID = ?", Id).Find(&getBook)
	return &getBook, results
}

func UpdateBook(Id int64, newName string, newAuthor string, newPublication string) {
	var book Book
	db.Where("ID = ?", Id).Find(&book)
	book.Name = newName
	book.Author = newAuthor
	book.Publication = newPublication
	db.Save(&book)
}

func DeleteBook(Id int64) Book {
	var book Book
	db.Where("ID = ?", Id).Delete(book)
	return book
}
