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
	"golang.org/x/exp/rand"
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

func GetMemberFees(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	memberId := vars["memberId"]

	ID, err := strconv.ParseUint(memberId, 10, 32)
	if err != nil {
		fmt.Println("Error while parsing")
	}

	// Fetch member details
	memberDetails, err := memberManager.GetMemberByID(uint(ID))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Member not found"))
		return
	}

	// Generate random number with 30
	overdueDays := rand.Intn(30)

	// Calculate total amount with penalties
	totalAmount := memberDetails.CalculateTotalAmount(overdueDays)

	// Prepare the response with fees breakdown
	response := map[string]interface{}{
		"member":        memberDetails,
		"outdated_fees": memberDetails.OutdatedFees,
		"overdue_days":  overdueDays,
		"total_amount":  totalAmount,
	}

	res, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
