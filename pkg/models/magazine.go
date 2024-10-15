package models

import (
	"github.com/jinzhu/gorm"
)

// Magazine datatype
type Magazine struct {
	gorm.Model
	Title       string `json:"title"`
	IssueNumber int    `json:"issue_number"`
	NumOfCopies int    `json:"num_of_copies"`
	Publisher   string `json:"publisher"`
	Year        int    `json:"year"`
}

// Interface implementing the magazine method signatures
type ManageMagazines interface {
	CreateMagazine(magazine *Magazine) *Magazine
	GetAllMagazines() []Magazine
	GetMagazineByID(magazineID uint) (*Magazine, error)
	UpdateMagazine(magazineID uint, magazine *Magazine)
}

// MagazineManager implements ManageAuthors for abstraction
type MagazineManager struct {
	DB *gorm.DB
}

// Create a new magazine
func (am *MagazineManager) CreateMagazine(magazine *Magazine) *Magazine {
	am.DB.Create(magazine)
	return magazine
}

// Retrieves all magazine from the database
func (am *MagazineManager) GetAllMagazines() []Magazine {
	var magazines []Magazine
	am.DB.Find(&magazines)
	return magazines
}

// Retrieves a Magazine by its ID
func (am *MagazineManager) GetMagazineByID(magazineID uint) (*Magazine, error) {
	var magazine Magazine
	if err := db.Where("id = ?", magazineID).Find(&magazine).Error; err != nil {
		return nil, err
	}
	return &magazine, nil
}

// updates an existing magazine in the database
func (am *MagazineManager) UpdateMagazine(magazineID uint, magazine *Magazine) {
	magazine.ID = magazineID
	am.DB.Save(magazine)
}

// Polymorphic methods implemented similar to book
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
