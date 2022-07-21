package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/goatking91/go-gin-study/practice2/internal/logo"
	_ "github.com/goatking91/go-gin-study/practice2/pkg/config"
	_ "github.com/goatking91/go-gin-study/practice2/pkg/db"
	"github.com/goatking91/go-gin-study/practice2/pkg/logger"
	_ "github.com/goatking91/go-gin-study/practice2/pkg/redis"
	"github.com/goatking91/go-gin-study/practice2/pkg/util"
)

func main() {
	env := &util.Env{EnvSource: &util.EnvGetter{}}
	gin.SetMode(env.GetString("SERVER_RUN_MODE"))

	r := gin.Default()

	logger.S.Info(logo.GenerateLogo())

	endPoint := fmt.Sprintf(":%v", env.GetString("SERVER_PORT"))
	readTimeout := env.GetDuration("SERVER_READ_TIMEOUT")
	writeTimeout := env.GetDuration("SERVER_WRITE_TIMEOUT")

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
