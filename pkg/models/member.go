package models

import (
	"github.com/jinzhu/gorm"
)

type Member struct {
	gorm.Model
	Person
	LoanRecords []LoanInformation `gorm:"foreignKey:MemberID" json:"loan_records"`
}

type LoanInformation struct {
	gorm.Model
	MemberID     uint   `json:"member_id"`
	LoanableID   uint   `json:"loanable_id"`
	LoanableType string `json:"loanable_type"`
	BorrowDate   string `json:"borrow_date"`
	ReturnDate   string `json:"return_date"`
	Member       Member `gorm:"foreignKey:MemberID" json:"member"`

	// These fields are used to store preloaded data for Book or Magazine
	LoanableBook     Book     `json:"loanable_book" gorm:"-"`
	LoanableMagazine Magazine `json:"loanable_magazine" gorm:"-"`
}

type MemberManager interface {
	CreateMember(member *Member) *Member
	GetMemberByID(memberID uint) (*Member, error)
}

// GormAuthorManager implements ManageAuthors using GORM
type GormMemberManager struct {
	DB *gorm.DB
}

// Create a new author
func (m *GormMemberManager) CreateMember(member *Member) *Member {
	m.DB.Create(member)
	return member
}

// Retrieves an member by their ID
func (m *GormMemberManager) GetMemberByID(memberID uint) (*Member, error) {
	var member Member
	if err := m.DB.Where("id = ?", memberID).Find(&member).Error; err != nil {
		return nil, err
	}
	return &member, nil
}

func (l *LoanInformation) CreateLoan(loanable Loanable) *LoanInformation {
	db.NewRecord(l)
	db.Create(&l)

	return l
}

// Get all loan information by MemberID
func GetLoansByID(memberID uint) []LoanInformation {
	var loans []LoanInformation
	// Preload both Book and Magazine based on LoanableType
	db.Where("member_id = ?", memberID).Preload("Member").Find(&loans)
	// Iterate over the loans and preload the appropriate loanable entity (Book or Magazine)
	for i := range loans {
		if loans[i].LoanableType == "book" {
			db.Where("id = ?", loans[i].LoanableID).First(&loans[i].LoanableBook)
		} else if loans[i].LoanableType == "magazine" {
			db.Where("id = ?", loans[i].LoanableID).First(&loans[i].LoanableMagazine)
		}
	}
	return loans
}
