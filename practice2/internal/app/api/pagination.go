package api

import (
	"github.com/gin-gonic/gin"
	"math"
	"strconv"
)

type Pagination struct {
	Offset      int   `json:"-"`
	Size        int   `json:"-"`
	CurrentPage int   `json:"currentPage"`
	TotalPage   int   `json:"totalPage"`
	TotalItems  int64 `json:"totalItems"`
}

type DataWithPagination struct {
	*Pagination
	Items interface{} `json:"items"`
}

const (
	defaultPage = 1
	defaultSize = 30
)

// Calc 페이지 계산하기
func (p *Pagination) Calc(total int64) *Pagination {
	p.Offset = p.Size * (p.CurrentPage - 1)
	p.TotalItems = total
	// calc maxpage
	d := float64(total) / float64(p.Size)
	p.TotalPage = int(math.Ceil(d))

	return p
}

// PaginationFromRequest request 에서 page, pageSize 가져오기
func PaginationFromRequest(c *gin.Context) *Pagination {
	p := &Pagination{}

	page := c.DefaultQuery("page", strconv.FormatInt(defaultPage, 10))
	pageSize := c.DefaultQuery("size", strconv.FormatInt(defaultSize, 10))

	p.CurrentPage, _ = strconv.Atoi(page)
	p.Size, _ = strconv.Atoi(pageSize)

	return p
}
