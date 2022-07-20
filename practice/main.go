package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func defaultHandler(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "default.html", gin.H{})
}

func setupRouter(engine *gin.Engine) {
	engine.LoadHTMLGlob("templates/**/*.html")
	engine.GET("/", defaultHandler)
}

func main() {
	r := gin.Default()
	setupRouter(r)
	err := r.Run(":3000")
	if err != nil {
		log.Fatalf("gin run error: %s", err)
	}
}
