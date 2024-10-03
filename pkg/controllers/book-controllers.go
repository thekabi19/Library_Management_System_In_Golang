package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

func GetBookByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) //access the book id in the request body
	bookId := vars["bookId"]

	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	bookDetails, _ := models.GetBookByID(ID)

	//response to postman
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	CreateBook := &models.Book{}
	utils.ParseBody(r, CreateBook)
	b := CreateBook.CreateBook()

	// Create a channel to communicate the completion of the Goroutine
	done := make(chan bool)

	// Use a Goroutine to send the notification asynchronously
	go utils.SendNotification(b.Name, done)

	// Respond to the client immediately
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

	// Wait for the notification Goroutine to finish (optional)
	go func() {
		<-done // Wait for the signal that the Goroutine has finished
		fmt.Println("Notification process completed.")
	}()
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	book := models.DeleteBook(ID)

	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
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

	bookDetails, db := models.GetBookByID(ID)
	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if updateBook.ISBN != "" {
		bookDetails.ISBN = updateBook.ISBN
	}
	if updateBook.NumberOfCopies != 0 {
		bookDetails.NumberOfCopies = updateBook.NumberOfCopies
	}
	if updateBook.AuthorID != 0 {
		bookDetails.AuthorID = updateBook.AuthorID
	}
	db.Save(&bookDetails)

	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// Get all books by an author
func GetBooksByAuthor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	authorId, err := strconv.Atoi(vars["authorId"])
	if err != nil {
		fmt.Println("error while parsing")
	}
	books := models.GetBooksByAuthor(uint(authorId))

	res, _ := json.Marshal(books)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetAuthor(w http.ResponseWriter, r *http.Request) {
	authors := models.GetAllAuthors()
	res, _ := json.Marshal(authors)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetAuthorByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	authorId := vars["authorId"]

	ID, err := strconv.ParseInt(authorId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	authorDetails, _ := models.GetAuthorByID(ID)
	res, _ := json.Marshal(authorDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateAuthor(w http.ResponseWriter, r *http.Request) {
	author := &models.Author{}
	utils.ParseBody(r, author)
	a := author.CreateAuthor()
	res, _ := json.Marshal(a)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteAuthor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	authorId := vars["authorId"]

	ID, err := strconv.ParseInt(authorId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	author := models.DeleteAuthor(ID)
	res, _ := json.Marshal(author)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// CreateBorrow handles the creation of a new borrow record with multiple books
func CreateBorrow(w http.ResponseWriter, r *http.Request) {
	borrow := &models.Borrow{}
	utils.ParseBody(r, borrow) // Parse the request body
	b := borrow.CreateBorrow() // Create a new borrow record
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// GetAllBorrows retrieves all borrow records with multiple books
func GetAllBorrows(w http.ResponseWriter, r *http.Request) {
	borrows := models.GetAllBorrows()
	res, _ := json.Marshal(borrows)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
