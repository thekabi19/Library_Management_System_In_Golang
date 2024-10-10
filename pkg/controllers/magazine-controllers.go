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

var magazineManager = &models.GormMagazineManager{DB: config.GetDB()} // Initialize GormAuthorManager

func GetAllMagazines(w http.ResponseWriter, r *http.Request) {
	magazines := magazineManager.GetAllMagazines()
	res, _ := json.Marshal(magazines)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetMagazineByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	magazineId := vars["magazineId"]

	ID, err := strconv.ParseInt(magazineId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}

	magazine, _ := magazineManager.GetMagazineByID(uint(ID))
	res, _ := json.Marshal(magazine)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateMagazine(w http.ResponseWriter, r *http.Request) {
	magazine := &models.Magazine{}
	utils.ParseBody(r, magazine)
	createdMagazine := magazineManager.CreateMagazine(magazine)
	res, _ := json.Marshal(createdMagazine)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateMagazine(w http.ResponseWriter, r *http.Request) {
	var updateMagazine = &models.Magazine{}
	utils.ParseBody(r, updateMagazine)
	vars := mux.Vars(r)
	magazineId := vars["magazineId"]
	ID, err := strconv.ParseInt(magazineId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}

	magazineDetails, err := magazineManager.GetMagazineByID(uint(ID)) // Use the interface method
	if err != nil {
		http.Error(w, "Magazine not found", http.StatusNotFound)
		return
	}
	if updateMagazine.Title != "" {
		magazineDetails.Title = updateMagazine.Title
	}
	if updateMagazine.NumOfCopies != 0 {
		magazineDetails.NumOfCopies = updateMagazine.NumOfCopies
	}
	if updateMagazine.IssueNumber != 0 {
		magazineDetails.IssueNumber = updateMagazine.IssueNumber
	}
	if updateMagazine.Year != 0 {
		magazineDetails.Year = updateMagazine.Year
	}
	if updateMagazine.Publisher != "" {
		magazineDetails.Publisher = updateMagazine.Publisher
	}
	magazineManager.UpdateMagazine(uint(ID), magazineDetails) // Use the interface method

	res, _ := json.Marshal(magazineDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
