package models

import (
	"github.com/jinzhu/gorm"
	"github.com/thekabi19/CSP3341_A2_code/pkg/config"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Title       string                `json:"title"`
	Year        int                   `json:"year"`
	AuthorID    uint                  `json:"author_id"`
	ISBN        string                `json:"isbn"`
	Publication string                `json:"publication"`
	NumOfCopies int                   `json:"num_of_copies"`
	Author      Author                `gorm:"foreignKey:AuthorID" json:"author"`
	LoanRecords []BookLoanInformation `gorm:"foreignKey:BookID" json:"loan_records"`
}

type Author struct {
	gorm.Model
	Name  string `json:"name"`
	Email string `json:"email"`
	Books []Book `gorm:"foreignKey:AuthorID" json:"books"`
}

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

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{}, &Author{}, &Member{}, &BookLoanInformation{}) // Automatically migrate the schema
}

func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b) // Create the book entry in the database
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookByID(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", Id).Find(&getBook)
	return &getBook, db
}

func DeleteBook(ID int64) Book {
	var book Book
	db.Where("ID=?", ID).Delete(book)
	return book
}

// Get all books by author
func GetBooksByAuthor(authorId uint) []Book {
	var books []Book
	db.Where("author_id = ?", authorId).Find(&books)
	return books
}

// Author CRUD operations
func (a *Author) CreateAuthor() *Author {
	db.NewRecord(a)
	db.Create(&a)
	return a
}

func GetAllAuthors() []Author {
	var Authors []Author
	db.Find(&Authors)
	return Authors
}

// Get an author by ID
func GetAuthorByID(ID int64) (*Author, *gorm.DB) {
	var author Author
	db := db.Preload("Books").Where("ID=?", ID).Find(&author)
	return &author, db
}

// Delete an author
func DeleteAuthor(ID int64) Author {
	var author Author
	db.Where("ID=?", ID).Delete(author)
	return author
}

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
