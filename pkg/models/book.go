package models

import (
	"github.com/jinzhu/gorm"
	"github.com/thekabi19/CSP3341_A2_code/pkg/config"
)

var db *gorm.DB

// Book represents a book entity in the library
type Book struct {
	gorm.Model
	Name           string `json:"name"`
	ISBN           string `json:"isbn"`
	Publication    string `json:"publication"`
	NumberOfCopies int    `json:"number_of_copies"`
	AuthorID       uint   `json:"author_id"` // Foreign key for author
}

// Author represents an author entity
type Author struct {
	gorm.Model
	Name  string `json:"name"`
	Books []Book `gorm:"foreignkey:AuthorID" json:"books"` // One-to-many relationship
}

// Borrow struct represents a borrowing transaction
type Borrow struct {
	gorm.Model
	BorrowerName string       `json:"borrower_name"`
	EmailAddress string       `json:"email_address"`
	BorrowDate   string       `json:"borrow_date"`
	ReturnDate   string       `json:"return_date"`
	LoanedBooks  []LoanedBook `json:"loaned_books"` // One-to-many relationship with LoanedBook
}

// LoanedBook represents a specific book loaned in a borrow transaction
type LoanedBook struct {
	gorm.Model
	BorrowID uint `json:"borrow_id"`                     // Foreign key to Borrow
	BookID   uint `json:"book_id"`                       // Link to Book model
	Book     Book `gorm:"foreignKey:BookID" json:"book"` // Book details
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{}, &Author{}, &Borrow{}, &LoanedBook{}) // Automatically migrate the schema
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

// CreateBorrow creates a new borrow record with multiple books
func (b *Borrow) CreateBorrow() *Borrow {
	db.NewRecord(b)
	db.Create(&b)
	for _, loanedBook := range b.LoanedBooks {
		loanedBook.BorrowID = b.ID // Auto-assign the borrow ID
		db.Create(&loanedBook)     // Create a new LoanedBook for each book
	}
	return b
}

// GetAllBorrows retrieves all borrowing records with loaned books
func GetAllBorrows() []Borrow {
	var borrows []Borrow
	db.Preload("LoanedBooks.Book").Find(&borrows) // Preload books for each borrow record
	return borrows
}
