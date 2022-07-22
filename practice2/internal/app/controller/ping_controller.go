package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"time"
)

type pong struct {
	URI      string `json:"uri"`
	Message  string `json:"message"`
	Hostname string `json:"hostname"`
	Datetime string `json:"datetime"`
}

// Ping
// @Summary ping
// @Schemes
// @Description 헬스체크를 위한 ping입니다.
// @Tags ping
// @Accept json
// @Produce json
// @Success 200 {object} pong
// @Router /ping [get]
func Ping(ctx *gin.Context) {
	hostname, _ := os.Hostname()
	ctx.JSON(http.StatusOK, pong{
		URI:      ctx.Request.RequestURI,
		Message:  "Pong! I am alive.",
		Hostname: hostname,
		Datetime: time.Now().Format("2006-01-02 15:04:05.000000 MST"),
	})
}
