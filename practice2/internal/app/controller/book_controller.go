package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"github.com/goatking91/go-gin-study/practice2/internal/app/api"
	"github.com/goatking91/go-gin-study/practice2/internal/app/model"
	"github.com/goatking91/go-gin-study/practice2/internal/app/service"
)

type bookController struct {
	bookService service.BookService
}

type BookController interface {
	CreateBook(*gin.Context)
	IndexBooks(*gin.Context)
	ShowBook(*gin.Context)
	DeleteBook(*gin.Context)
	UpdateBook(*gin.Context)
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
// @Success 201 {object} api.SuccessRes{data=model.Book}
// @Failure 400 {object} api.ErrorRes{error=api.Err}
// @Router /books [post]
func (b bookController) CreateBook(ctx *gin.Context) {
	var book model.Book

	if err := ctx.ShouldBindJSON(&book); err != nil {
		_ = ctx.Error(err)
		return
	}

	book, err := b.bookService.CreateBook(book)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	api.Response(ctx, http.StatusCreated, book)
}

// IndexBooks
// @Summary 책 리스트
// @Schemes
// @Description 등록된 책의 list를 보는 API입니다.
// @Tags books
// @Accept json
// @Produce json
// @Success 200 {object} api.SuccessRes{data=api.Pagination{items=[]model.Book}}
// @Router /books [get]
func (b bookController) IndexBooks(ctx *gin.Context) {
	p := api.PaginationFromRequest(ctx)
	books, err := b.bookService.GetBooks(p)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	data := api.DataWithPagination{
		Pagination: p,
		Items:      books,
	}

	api.Response(ctx, http.StatusOK, data)
}

// ShowBook
// @Summary 책 상세 조회
// @Schemes
// @Description 책의 정보를 상세하게 보는 API입니다.
// @Tags books
// @Accept json
// @Produce json
// @Param uid path string true "book id"
// @Success 200 {object} api.SuccessRes{data=model.Book}
// @Failure 404 {object} api.ErrorRes{error=api.Err}
// @Router /books/{uid} [get]
func (b bookController) ShowBook(ctx *gin.Context) {
	uid := ctx.Param("uid")
	book, err := b.bookService.GetBook(uid)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	api.Response(ctx, http.StatusOK, book)
}

// DeleteBook
// @Summary 책 삭제
// @Schemes
// @Description 책의 정보를 삭제하는 API입니다.
// @Tags books
// @Accept json
// @Produce json
// @Param uid path string true "book id"
// @Success 204
// @Failure 404 {object} api.ErrorRes{error=api.Err}
// @Router /books/{uid} [delete]
func (b bookController) DeleteBook(ctx *gin.Context) {
	uid := ctx.Param("uid")
	book, err := b.bookService.DeleteBook(uid)

	if err != nil {
		_ = ctx.Error(err)
		return
	}

	api.Response(ctx, http.StatusNoContent, book)
}

// UpdateBook
// @Summary 책 업데이트
// @Schemes
// @Description 책의 정보를 수정하는 API입니다.
// @Tags books
// @Accept json
// @Produce json
// @Param uid path string true "book id"
// @Param data body model.Book true "The input book struct"
// @Success 200 {object} api.SuccessRes{data=model.Book}
// @Failure 400 {object} api.ErrorRes{error=api.Err}
// @Failure 404 {object} api.ErrorRes{error=api.Err}
// @Router /books/{uid} [put]
func (b bookController) UpdateBook(ctx *gin.Context) {
	uid := ctx.Param("uid")
	var body model.Book
	if err := ctx.ShouldBindJSON(&body); err != nil {
		_ = ctx.Error(err)
		return
	}

	book, err := b.bookService.UpdateBook(uid, body)

	if err != nil {
		_ = ctx.Error(err)
		return
	}

	api.Response(ctx, http.StatusOK, book)
}
