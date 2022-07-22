package service

import (
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/goatking91/go-gin-study/practice2/internal/app/model"
	"github.com/goatking91/go-gin-study/practice2/pkg/db"
)

type bookService struct {
	db *gorm.DB
}

type BookService interface {
	CreateBook(book model.Book) (model.Book, error)
	GetBooks() ([]model.Book, error)
}

func NewBookService(db *gorm.DB) BookService {
	return bookService{db: db}
}

func (b bookService) CreateBook(book model.Book) (model.Book, error) {
	book.UID = uuid.New().String()
	err := db.DB.Create(&book).Error
	return book, err
}

func (b bookService) GetBooks() ([]model.Book, error) {
	var books []model.Book
	result := db.DB.Find(&books)
	return books, result.Error
}
