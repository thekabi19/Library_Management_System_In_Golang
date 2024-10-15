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

var authorManager = &models.AuthorManager{DB: config.GetDB()} // Initialize AuthorManager

// Get all books by an author
func GetBooksByAuthor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	authorId, err := strconv.Atoi(vars["authorId"]) //convert the authorID string to integer using Atoi function (Open My Mind, 2023)
	if err != nil {                                 //handle parsing errors
		fmt.Println("error while parsing")
	}
	books := models.GetBooksByAuthor(uint(authorId)) //convert ID to unsigned Integer incase

	res, _ := json.Marshal(books) //marshal to JSON format
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK) //send it back to the endpoint
	w.Write(res)
}

// GetAllAuthors retrieves all authors
func GetAllAuthors(w http.ResponseWriter, r *http.Request) {
	authors := authorManager.GetAllAuthors()
	res, _ := json.Marshal(authors)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// GetAuthorByID retrieves an author by their ID
func GetAuthorByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	authorIDStr := vars["authorId"]                         //Get the authorID which is string
	authorID, err := strconv.ParseUint(authorIDStr, 10, 32) //convert it to Unsigned Int with base size 10 and bit size 32 with uint32 (IncludeHelp, 2021)
	if err != nil {
		fmt.Println("Error while parsing author ID")
	}

	author, err := authorManager.GetAuthorByID(uint(authorID))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	res, _ := json.Marshal(author)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// CreateAuthor creates a new author
func CreateAuthor(w http.ResponseWriter, r *http.Request) {
	newAuthor := &models.Author{}
	utils.ParseBody(r, newAuthor) //parses the http request body into newAuthor struct
	createdAuthor := authorManager.CreateAuthor(newAuthor)
	res, _ := json.Marshal(createdAuthor)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// DeleteAuthor deletes an author by their ID
func DeleteAuthor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	authorIDStr := vars["authorId"]
	authorID, err := strconv.ParseUint(authorIDStr, 10, 32) //convert it to Unsigned Int with base size 10 and bit size 32 with uint32 (IncludeHelp, 2021)
	if err != nil {
		fmt.Println("Error while parsing author ID")
	}

	author := authorManager.DeleteAuthor(uint(authorID))
	res, _ := json.Marshal(author)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
