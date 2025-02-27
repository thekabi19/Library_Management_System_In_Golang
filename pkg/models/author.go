package models

import (
	"github.com/jinzhu/gorm"
)

// Person struct to hold common fields for Author and Member
type Person struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Author struct embedding Person
type Author struct {
	gorm.Model
	Person
	Books []Book `gorm:"foreignKey:AuthorID" json:"books"`
}

type AuthorManager interface {
	CreateAuthor(author *Author) *Author
	GetAllAuthors() []Author
	GetAuthorByID(authorID uint) (*Author, error)
	DeleteAuthor(authorID uint) *Author
}

// GormAuthorManager implements ManageAuthors using GORM
type GormAuthorManager struct {
	DB *gorm.DB
}

// Get all books by author
func GetBooksByAuthor(authorId uint) []Book {
	var books []Book
	db.Where("author_id = ?", authorId).Find(&books)
	return books
}

// Create a new author
func (am *GormAuthorManager) CreateAuthor(author *Author) *Author {
	am.DB.Create(author)
	return author
}

// Retrieves all authors from the database
func (am *GormAuthorManager) GetAllAuthors() []Author {
	var authors []Author
	am.DB.Find(&authors)
	return authors
}

// Retrieves an author by their ID
func (am *GormAuthorManager) GetAuthorByID(authorID uint) (*Author, error) {
	var author Author
	if err := am.DB.Preload("Books").Where("id = ?", authorID).Find(&author).Error; err != nil {
		return nil, err
	}
	return &author, nil
}

// Removes an author from the database by their ID
func (am *GormAuthorManager) DeleteAuthor(authorID uint) *Author {
	var author Author
	am.DB.Where("id = ?", authorID).Delete(&author)
	return &author
}
