package models

import (
	"github.com/jinzhu/gorm"
)

type Member struct {
	gorm.Model
	Name         string                `json:"name"`
	EmailAddress string                `json:"email_address"`
	LoanRecords  []BookLoanInformation `gorm:"foreignKey:MemberID" json:"loan_records"`
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

/* GormAuthorManager implements ManageAuthors using GORM
type GormMemberManager struct {
	DB *gorm.DB
}*/

func (m *Member) CreateMember() *Member {
	db.NewRecord(m)
	db.Create(&m)
	return m
}

func GetMemberByID(Id int64) (*Member, *gorm.DB) {
	var getMember Member
	db := db.Where("ID=?", Id).Find(&getMember)
	return &getMember, db
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
