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

// ManageBooks interface for managing book-related operations
type ManageBooks interface {
	CreateBook(book *Book)
	UpdateBook(bookID uint, book *Book)
	DeleteBook(bookID uint)
	GetBookByID(bookID uint) (*Book, error)
}

type GormBookManager struct {
	DB *gorm.DB // Changed from 'db' to 'DB' to make it exported
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{}, &Author{}, &Member{}, &BookLoanInformation{}) // Automatically migrate the schema
}

// AddBook adds a new book to the database
func (bm *GormBookManager) CreateBook(book *Book) {
	bm.DB.Create(book)
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

// UpdateBook updates an existing book in the database
func (bm *GormBookManager) UpdateBook(bookID uint, book *Book) {
	book.ID = bookID
	bm.DB.Save(book)
}

// GetBookByID retrieves a book by its ID
func (bm *GormBookManager) GetBookByID(bookID uint) (*Book, error) {
	var book Book
	if err := bm.DB.Where("id = ?", bookID).Find(&book).Error; err != nil {
		return nil, err
	}
	return &book, nil
}

// DeleteBook removes a book from the database by its ID
func (bm *GormBookManager) DeleteBook(bookID uint) {
	bm.DB.Delete(&Book{}, bookID)
}
