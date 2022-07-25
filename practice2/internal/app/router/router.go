package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/goatking91/go-gin-study/practice2/internal/app/controller"
	"github.com/goatking91/go-gin-study/practice2/internal/app/middleware"
	"github.com/goatking91/go-gin-study/practice2/internal/app/service"
	"github.com/goatking91/go-gin-study/practice2/pkg/db"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.HandleMethodNotAllowed = true

	r.Use(gin.Logger())

	r.Use(middleware.CorsMiddleware())
	r.Use(middleware.HttpLogger())
	r.Use(middleware.Error())

	//r.NoMethod(middleware.Error())
	//r.NoRoute()

	createRoutes(r)

	return r
}

func createRoutes(r *gin.Engine) {
	bookService := service.NewBookService(db.DB)
	bookController := controller.NewBookController(bookService)

	v1 := r.Group("/v1")

	v1.GET("/ping", controller.Ping)

	books := v1.Group("/books")
	books.POST("", bookController.CreateBook)
	books.GET("", bookController.IndexBooks)
	books.GET("/:uid", bookController.ShowBook)
	books.DELETE("/:uid", bookController.DeleteBook)
	books.PUT("/:uid", bookController.UpdateBook)

	v1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
