package models

import (
	"github.com/jinzhu/gorm"
)

type Member struct {
	gorm.Model
	Person
	LoanRecords []BookLoanInformation `gorm:"foreignKey:MemberID" json:"loan_records"`
}

type BookLoanInformation struct {
	gorm.Model
	MemberID   uint   `json:"member_id"`
	BookID     uint   `json:"book_id"`
	BorrowDate string `json:"borrow_date"`
	ReturnDate string `json:"return_date"`
	Member     Member `gorm:"foreignKey:MemberID" json:"member"`
	Book       Book   `gorm:"foreignKey:BookID" json:"book"`
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
	if err := m.DB.Preload("LoanRecords").Where("id = ?", memberID).Find(&member).Error; err != nil {
		return nil, err
	}
	return &member, nil
}

func (l *BookLoanInformation) CreateLoan() *BookLoanInformation {
	db.NewRecord(l)
	db.Create(&l)
	return l
}

// Get all loan information by MemberID
func GetLoansByMemberID(memberID uint) []BookLoanInformation {
	var loans []BookLoanInformation
	db.Where("member_id = ?", memberID).Preload("Member").Preload("Book").Find(&loans)
	return loans
}
