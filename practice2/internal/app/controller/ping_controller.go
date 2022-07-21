package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"time"
)

func Ping(ctx *gin.Context) {
	hostname, _ := os.Hostname()
	ctx.JSON(http.StatusOK, gin.H{
		"uri":      ctx.Request.RequestURI,
		"message":  "Pong! I am alive.",
		"hostname": hostname,
		"datetime": time.Now().Format("2006-01-02 15:04:05.000000 MST"),
	})
}
