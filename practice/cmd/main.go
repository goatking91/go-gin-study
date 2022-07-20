package main

import (
	"github.com/gin-gonic/gin"
	"github.com/goatking91/go-gin-study/practice/internal/books"
	"github.com/goatking91/go-gin-study/practice/pkg/db"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("./pkg/envs/.env")
	viper.ReadInConfig()

	port := viper.Get("PORT").(string)
	dbUrl := viper.Get("DB_URL").(string)

	r := gin.Default()
	h := db.Init(dbUrl)

	books.RegisterRoutes(r, h)

	r.Run(port)
}
