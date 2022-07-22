package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"github.com/goatking91/go-gin-study/practice2/internal/app/model"
	"github.com/goatking91/go-gin-study/practice2/internal/app/service"
)

type bookController struct {
	bookService service.BookService
}

type BookController interface {
	CreateBook(*gin.Context)
	GetBooks(*gin.Context)
}

func NewBookController(bookService service.BookService) BookController {
	return bookController{
		bookService: bookService,
	}
}

// CreateBook
// @Summary 책 추가
// @Schemes
// @Description 책을 추가하는 API입니다.
// @Tags books
// @Accept json
// @Produce json
// @Param data body model.Book true "The input book struct"
// @Success 200 {object} model.Book
// @Router /books [post]
func (b bookController) CreateBook(ctx *gin.Context) {
	var book model.Book
	if err := ctx.BindJSON(&book); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	book, err := b.bookService.CreateBook(book)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	ctx.IndentedJSON(http.StatusCreated, gin.H{"data": book})
}

// GetBooks
// @Summary 책 리스트
// @Schemes
// @Description 등록된 책의 list를 보는 API입니다.
// @Tags books
// @Accept json
// @Produce json
// @Success 200 {object} []model.Book
// @Router /books [get]
func (b bookController) GetBooks(ctx *gin.Context) {
	books, err := b.bookService.GetBooks()
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	ctx.IndentedJSON(http.StatusCreated, gin.H{"data": books})
}
