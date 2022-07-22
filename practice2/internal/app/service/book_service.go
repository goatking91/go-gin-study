package service

import (
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/goatking91/go-gin-study/practice2/internal/app/model"
)

type bookService struct {
	db *gorm.DB
}

type BookService interface {
	CreateBook(book model.Book) (model.Book, error)
	GetBooks() ([]model.Book, error)
	GetBook(uid string) (model.Book, error)
	DeleteBook(uid string) (model.Book, error)
	UpdateBook(uid string, body model.Book) (book model.Book, err error)
}

func NewBookService(db *gorm.DB) BookService {
	return bookService{db: db}
}

func (b bookService) CreateBook(book model.Book) (model.Book, error) {
	book.UID = uuid.New().String()
	err := b.db.Create(&book).Error
	return book, err
}

func (b bookService) GetBooks() ([]model.Book, error) {
	var books []model.Book
	result := b.db.Find(&books)
	return books, result.Error
}

func (b bookService) GetBook(uid string) (model.Book, error) {
	var book model.Book
	result := b.db.Where("uid = ?", uid).First(&book)
	return book, result.Error
}

func (b bookService) DeleteBook(uid string) (model.Book, error) {
	var book model.Book
	result := b.db.Where("uid = ?", uid).First(&book)

	if result.Error != nil {
		return book, result.Error
	}

	result = b.db.Delete(&book)
	return book, result.Error
}

func (b bookService) UpdateBook(uid string, body model.Book) (book model.Book, err error) {
	result := b.db.Where("uid = ?", uid).First(&book)

	if result.Error != nil {
		return book, result.Error
	}

	book.Title = body.Title
	book.Author = body.Author
	book.Description = body.Description

	result = b.db.Save(&book)
	err = result.Error
	return
}
