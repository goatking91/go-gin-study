package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/goatking91/go-gin-study/practice/pkg/models"
)

func Init(url string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(url), &gorm.Config{SkipDefaultTransaction: true})

	if err != nil {
		return db, err
	}

	db.AutoMigrate(&models.Book{})

	return db, nil
}
