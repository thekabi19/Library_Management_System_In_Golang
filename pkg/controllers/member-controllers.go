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

var memberManager = &models.GormMemberManager{DB: config.GetDB()} // Initialize GormMemberManager

func CreateMember(w http.ResponseWriter, r *http.Request) {
	newMember := &models.Member{}
	utils.ParseBody(r, &newMember)
	CreatedMember := memberManager.CreateMember(newMember)

	res, _ := json.Marshal(CreatedMember)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetMemberByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	memberId := vars["memberId"]

	ID, err := strconv.ParseUint(memberId, 10, 32)
	if err != nil {
		fmt.Println("error while parsing")
	}
	memberDetails, _ := memberManager.GetMemberByID(uint(ID))
	res, _ := json.Marshal(memberDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBookLoanInformation(w http.ResponseWriter, r *http.Request) {
	var newLoan models.BookLoanInformation
	utils.ParseBody(r, &newLoan)

	loan := newLoan.CreateLoan()

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
	loans := models.GetLoansByMemberID(uint(memberID))

	// Convert the result to JSON and send response
	res, _ := json.Marshal(loans)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
