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

var authorManager = &models.GormAuthorManager{DB: config.GetDB()} // Initialize GormBookManager

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

// GetAllAuthors retrieves all authors
func GetAllAuthors(w http.ResponseWriter, r *http.Request) {
	authors := authorManager.GetAllAuthors()
	res, _ := json.Marshal(authors)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// GetAuthorByID retrieves an author by their ID
func GetAuthorByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	authorIDStr := vars["authorId"]
	authorID, err := strconv.ParseUint(authorIDStr, 10, 32)
	if err != nil {
		fmt.Println("Error while parsing author ID")
	}

	author, err := authorManager.GetAuthorByID(uint(authorID))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	res, _ := json.Marshal(author)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// CreateAuthor creates a new author
func CreateAuthor(w http.ResponseWriter, r *http.Request) {
	newAuthor := &models.Author{}
	utils.ParseBody(r, newAuthor)
	createdAuthor := authorManager.CreateAuthor(newAuthor)
	res, _ := json.Marshal(createdAuthor)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// DeleteAuthor deletes an author by their ID
func DeleteAuthor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	authorIDStr := vars["authorId"]
	authorID, err := strconv.ParseUint(authorIDStr, 10, 32)
	if err != nil {
		fmt.Println("Error while parsing author ID")
	}

	author := authorManager.DeleteAuthor(uint(authorID))
	res, _ := json.Marshal(author)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
