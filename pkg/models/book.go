package models

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/thekabi19/CSP3341_A2_code/pkg/config"
)

var db *gorm.DB

// Loanable interface has method signatures common for any loanble items such as the book and magazine
type Loanable interface {
	GetID() uint
	GetTitle() string
	GetNumOfCopies() int
	DecrementCopies()
}

// Book Struct defines the Book datatype
type Book struct {
	gorm.Model
	Title       string `json:"title"`
	Year        int    `json:"year"`
	AuthorID    uint   `json:"author_id"`
	ISBN        string `json:"isbn"`
	Publication string `json:"publication"`
	NumOfCopies int    `json:"num_of_copies"`
	Author      Author `gorm:"foreignKey:AuthorID" json:"author"`
	//LoanRecords []BookLoanInformation `gorm:"foreignKey:BookID" json:"loan_records"`
}

// ManageBooks interface for managing book-related operations
type ManageBooks interface {
	CreateBook(book *Book)
	UpdateBook(bookID uint, book *Book)
	DeleteBook(bookID uint)
	GetBookByID(bookID uint) (*Book, error)
}

var ErrBookNotFound = errors.New("Book with the entered ID not found")

// A book manager struct with pointer to the mysql database object
type BookManager struct {
	DB *gorm.DB
}

// Intializes the tables in the library database
func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{}, &Magazine{}, &Author{}, &Member{}, &LoanInformation{}) // Automatically migrate the schema
}

// AddBook adds a new book to the database
func (bm *BookManager) CreateBook(book *Book) {
	bm.DB.Create(book)
}

// GetAllBooks get all books from the database
func GetAllBooks() []Book {
	var Books []Book
	db.Preload("Author").Find(&Books)
	return Books
}

// UpdateBook updates an existing book in the database
func (bm *BookManager) UpdateBook(bookID uint, book *Book) {
	book.ID = bookID
	bm.DB.Save(book)
}

// GetBookByID retrieves a book by its bookID
func (bm *BookManager) GetBookByID(bookID uint) (*Book, error) {
	var book Book
	bookFound := bm.DB.Where("id = ?", bookID).Preload("Author").Find(&book) //preload author as we need the authur also for the book
	//Error checking if the book is found or not
	if bookFound.RowsAffected == 0 {
		return nil, ErrBookNotFound
	} else if bookFound.Error != nil {
		return nil, bookFound.Error
	}

	return &book, nil
}

// DeleteBook removes a book from the database by its ID
func (bm *BookManager) DeleteBook(bookID uint) {
	bm.DB.Delete(&Book{}, bookID)
}

// Polymorphic methods for book
func (b *Book) GetID() uint {
	return b.ID
}

func (b *Book) GetTitle() string {
	return b.Title
}

func (b *Book) GetNumOfCopies() int {
	return b.NumOfCopies
}

func (b *Book) DecrementCopies() {
	if b.NumOfCopies > 0 {
		b.NumOfCopies--
	}
}
