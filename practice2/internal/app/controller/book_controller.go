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
	Create(*gin.Context)
}

func NewBookController(bookService service.BookService) BookController {
	return bookController{
		bookService: bookService,
	}
}

func (b bookController) Create(ctx *gin.Context) {
	var book model.Book
	if err := ctx.BindJSON(&book); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	book, err := b.bookService.Create(book)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	ctx.IndentedJSON(http.StatusCreated, gin.H{"data": book})
}
