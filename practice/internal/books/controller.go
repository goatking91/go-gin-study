package books

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	routes := r.Group("/books")
	routes.POST("/", h.AddBook)
	routes.GET("/", h.IndexBooks)
	routes.GET("/:id", h.ShowBook)
	routes.PUT("/:id", h.UpdateBook)
	routes.DELETE("/:id", h.DeleteBook)
}
