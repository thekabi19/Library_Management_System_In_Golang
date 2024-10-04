package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/thekabi19/CSP3341_A2_code/pkg/config"
	"github.com/thekabi19/CSP3341_A2_code/pkg/models"
	"github.com/thekabi19/CSP3341_A2_code/pkg/utils"
)

var NewBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

var bookManager = &models.GormBookManager{DB: config.GetDB()} // Initialize GormBookManager

// GetBookByID retrieves a book by its ID
func GetBookByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) // Access the book ID in the request body
	bookId := vars["bookId"]

	ID, err := strconv.ParseUint(bookId, 10, 0)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	bookDetails, err := bookManager.GetBookByID(uint(ID)) // Use the interface method
	if err != nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// CreateBook adds a new book to the database
func CreateBook(w http.ResponseWriter, r *http.Request) {
	newBook := &models.Book{}
	utils.ParseBody(r, newBook)
	bookManager.CreateBook(newBook) // Use the interface method

	// Create a channel to communicate the completion of the Goroutine
	done := make(chan bool)

	// Use a Goroutine to send the notification asynchronously
	go utils.SendNotification(newBook.Title, done)

	// Respond to the client immediately
	res, _ := json.Marshal(newBook)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

	// Wait for the notification Goroutine to finish (optional)
	go func() {
		<-done // Wait for the signal that the Goroutine has finished
		fmt.Println("Notification process completed.")
	}()
}

// DeleteBook removes a book from the database by its ID
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseUint(bookId, 10, 0)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	bookManager.DeleteBook(uint(ID)) // Use the interface method

	w.WriteHeader(http.StatusNoContent) // No content returned on delete
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var updateBook = &models.Book{}
	utils.ParseBody(r, updateBook)
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}

	bookDetails, err := bookManager.GetBookByID(uint(ID)) // Use the interface method
	if err != nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}
	if updateBook.Title != "" {
		bookDetails.Title = updateBook.Title
	}
	if updateBook.ISBN != "" {
		bookDetails.ISBN = updateBook.ISBN
	}
	if updateBook.NumOfCopies != 0 {
		bookDetails.NumOfCopies = updateBook.NumOfCopies
	}
	if updateBook.AuthorID != 0 {
		bookDetails.AuthorID = updateBook.AuthorID
	}
	if updateBook.Year != 0 {
		bookDetails.Year = updateBook.Year
	}
	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}
	bookManager.UpdateBook(uint(ID), bookDetails) // Use the interface method

	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
