package models

import (
	"github.com/jinzhu/gorm"
)

type Magazine struct {
	gorm.Model
	Title       string `json:"title"`
	IssueNumber int    `json:"issue_number"`
	NumOfCopies int    `json:"num_of_copies"`
	Publisher   string `json:"publisher"`
	Year        int    `json:"year"`
}

type MagazineManager interface {
	CreateMagazine(magazine *Magazine) *Magazine
	GetAllMagazines() []Magazine
	GetMagazineByID(magazineID uint) (*Magazine, error)
	UpdateMagazine(magazineID uint, magazine *Magazine)
}

// GormMagazineManager implements ManageAuthors using GORM
type GormMagazineManager struct {
	DB *gorm.DB
}

// Create a new author
func (am *GormMagazineManager) CreateMagazine(magazine *Magazine) *Magazine {
	am.DB.Create(magazine)
	return magazine
}

// Retrieves all authors from the database
func (am *GormMagazineManager) GetAllMagazines() []Magazine {
	var magazines []Magazine
	am.DB.Find(&magazines)
	return magazines
}

// Retrieves a Magazine by their ID
func (am *GormMagazineManager) GetMagazineByID(magazineID uint) (*Magazine, error) {
	var magazine Magazine
	if err := db.Where("id = ?", magazineID).Find(&magazine).Error; err != nil {
		return nil, err
	}
	return &magazine, nil
}

// UpdateBook updates an existing book in the database
func (am *GormMagazineManager) UpdateMagazine(magazineID uint, magazine *Magazine) {
	magazine.ID = magazineID
	am.DB.Save(magazine)
}

func (m *Magazine) GetID() uint {
	return m.ID
}

func (m *Magazine) GetTitle() string {
	return m.Title
}

func (m *Magazine) GetNumOfCopies() int {
	return m.NumOfCopies
}

func (m *Magazine) DecrementCopies() {
	if m.NumOfCopies > 0 {
		m.NumOfCopies--
	}
}
