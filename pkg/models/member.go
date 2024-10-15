package models

import (
	"github.com/jinzhu/gorm"
)

// Member Datatype
type Member struct {
	gorm.Model
	Person
	OutdatedFees float32           `json:"outdated_fees"`
	LoanRecords  []LoanInformation `gorm:"foreignKey:MemberID" json:"loan_records"`
}

// Loan Information datatype
type LoanInformation struct {
	gorm.Model
	MemberID     uint   `json:"member_id"`
	LoanableID   uint   `json:"loanable_id"`   //stores either book ID or magazine ID
	LoanableType string `json:"loanable_type"` //specifies whether the item stored it book or magazine
	BorrowDate   string `json:"borrow_date"`
	ReturnDate   string `json:"return_date"`
	Member       Member `gorm:"foreignKey:MemberID" json:"member"`

	// These fields stores temporary data for Book or Magazine when called for all loaned items by each member
	LoanableBook     Book     `json:"loanable_book" gorm:"-"` //"-" is used to ensure this is not stored in the database
	LoanableMagazine Magazine `json:"loanable_magazine" gorm:"-"`
}

// Implements abstracted methods
type ManageMember interface {
	CreateMember(member *Member) *Member
	GetMemberByID(memberID uint) (*Member, error)
	CalculateTotalAmount(overdueDays int) float32
}

// AuthorManager implements ManageMember
type MemberManager struct {
	DB *gorm.DB
}

// Create a new member
func (m *MemberManager) CreateMember(member *Member) *Member {
	m.DB.Create(member)
	return member
}

// Retrieves an member by their ID
func (m *MemberManager) GetMemberByID(memberID uint) (*Member, error) {
	var member Member
	if err := m.DB.Where("id = ?", memberID).Find(&member).Error; err != nil {
		return nil, err
	}
	return &member, nil
}

// creates lones
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
	// Iterate over the loans and preload the appropriate loanable item
	for i := range loans {
		if loans[i].LoanableType == "book" { //if the loaned item is book
			db.Where("id = ?", loans[i].LoanableID).First(&loans[i].LoanableBook)
		} else if loans[i].LoanableType == "magazine" { //if the loaned item is magazine
			db.Where("id = ?", loans[i].LoanableID).First(&loans[i].LoanableMagazine)
		}
	}
	return loans
}

// Calculates the outdated fees for each member
func (m *Member) CalculateTotalAmount(overdueDays int) float32 {
	var totalAmount float32
	totalAmount = m.OutdatedFees

	if overdueDays > 3 && overdueDays <= 7 {
		totalAmount += totalAmount * 0.10 // 10% penalty for overdue between 3 to 7 days
	} else if overdueDays > 7 && overdueDays <= 30 {
		totalAmount += totalAmount * 0.30 // 30% penalty for overdue between 7 to 30 days
	} else if overdueDays > 30 {
		totalAmount += totalAmount * 0.50 // 50% penalty if overdue more than 30 days
	}

	return totalAmount
}
