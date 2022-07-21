package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/goatking91/go-gin-study/practice2/internal/app/controller"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.HandleMethodNotAllowed = true

	r.Use(gin.Logger())

	createRoutes(r)

	return r
}

func createRoutes(r *gin.Engine) {
	v1 := r.Group("/v1")

	v1.GET("/ping", controller.Ping)

	v1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
