package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/goatking91/go-gin-study/practice2/pkg/config"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	config.InitConfig()

	gin.SetMode(os.Getenv("SERVER_RUN_MODE"))

	r := gin.Default()

	endPoint := fmt.Sprintf(":%v", os.Getenv("SERVER_PORT"))
	readTimeout, _ := time.ParseDuration(os.Getenv("SERVER_READ_TIMEOUT"))
	writeTimeout, _ := time.ParseDuration(os.Getenv("SERVER_READ_TIMEOUT"))

	server := &http.Server{
		Addr:         endPoint,
		Handler:      r,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
	}

	log.Printf("Starting http server listening:(%s) timeout r:%v w:%v", endPoint, readTimeout, writeTimeout)

	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("Fail start http server. %v", err)
	}
}
