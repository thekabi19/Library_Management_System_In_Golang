package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/thekabi19/CSP3341_A2_code/pkg/models"
	"github.com/thekabi19/CSP3341_A2_code/pkg/utils"
)

func CreateLoanInformation(w http.ResponseWriter, r *http.Request) {
	var newLoan models.LoanInformation
	utils.ParseBody(r, &newLoan)

	// Determine if it's a book or magazine
	var loanable models.Loanable
	if newLoan.LoanableType == "book" {
		loanable, _ = bookManager.GetBookByID(newLoan.LoanableID)
	} else if newLoan.LoanableType == "magazine" {
		loanable, _ = magazineManager.GetMagazineByID(newLoan.LoanableID)
	}

	if loanable == nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Item not found"))
		return
	}

	// Check if there are copies available
	if loanable.GetNumOfCopies() <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("No copies available for borrowing"))
		return
	}
	// Decrement the number of available copies
	loanable.DecrementCopies()

	// Save the changes back to the database based on type
	if newLoan.LoanableType == "book" {
		bookManager.UpdateBook(newLoan.LoanableID, loanable.(*models.Book)) // Cast to *models.Book
	} else if newLoan.LoanableType == "magazine" {
		magazineManager.UpdateMagazine(newLoan.LoanableID, loanable.(*models.Magazine)) // Cast to *models.Magazine
	}

	//Add the return time automatically (Sai, 2023)
	now := time.Now()
	newLoan.BorrowDate = now.Format("2006-01-02")
	// Set the return date to 10 days after the borrow date
	newLoan.ReturnDate = now.AddDate(0, 0, 10).Format("2006-01-02")

	loan := newLoan.CreateLoan(loanable)

	res, _ := json.Marshal(loan)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetLoansForMember(w http.ResponseWriter, r *http.Request) {
	// Get the memberID from the URL or request
	vars := mux.Vars(r)
	memberID, err := strconv.Atoi(vars["memberID"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid member ID"))
		return
	}

	// Fetch all loan records for this member
	loans := models.GetLoansByID(uint(memberID))

	// Convert the result to JSON and send response
	res, _ := json.Marshal(loans)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
