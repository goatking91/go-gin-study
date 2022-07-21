package main

import (
	"fmt"
	"github.com/goatking91/go-gin-study/practice2/internal/logo"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/goatking91/go-gin-study/practice2/pkg/config"
	"github.com/goatking91/go-gin-study/practice2/pkg/logger"
)

func main() {
	gin.SetMode(os.Getenv("SERVER_RUN_MODE"))

	r := gin.Default()

	logger.S.Info(logo.GenerateLogo())

	endPoint := fmt.Sprintf(":%v", os.Getenv("SERVER_PORT"))
	readTimeout, _ := time.ParseDuration(os.Getenv("SERVER_READ_TIMEOUT"))
	writeTimeout, _ := time.ParseDuration(os.Getenv("SERVER_READ_TIMEOUT"))

	server := &http.Server{
		Addr:         endPoint,
		Handler:      r,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
	}

	logger.S.Infof("Starting http server listening:(%s) timeout r:%v w:%v", endPoint, readTimeout, writeTimeout)

	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("Fail start http server. %v", err)
	}
}
