package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/goatking91/go-gin-study/practice2/docs"
	"github.com/goatking91/go-gin-study/practice2/internal/app/router"
	"github.com/goatking91/go-gin-study/practice2/internal/logo"
	_ "github.com/goatking91/go-gin-study/practice2/pkg/config"
	"github.com/goatking91/go-gin-study/practice2/pkg/db"
	"github.com/goatking91/go-gin-study/practice2/pkg/logger"
	"github.com/goatking91/go-gin-study/practice2/pkg/redis"
	_ "github.com/goatking91/go-gin-study/practice2/pkg/redis"
	"github.com/goatking91/go-gin-study/practice2/pkg/util"
)

func init() {
	logger.InitLogger()
	db.InitDb()
	db.InitMigrate()
	redis.InitRedis()
}

// @title Practice2 API
// @version 1.0
func main() {
	env := &util.Env{EnvSource: &util.EnvGetter{}}
	gin.SetMode(env.GetString("SERVER_RUN_MODE"))

	docs.SwaggerInfo.Description = env.GetString("SWAGGER_DESCRIPTION")
	docs.SwaggerInfo.Host = env.GetString("SWAGGER_HOST")
	docs.SwaggerInfo.BasePath = env.GetString("SWAGGER_BASE_PATH")
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	r := router.InitRouter()

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
