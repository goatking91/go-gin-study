package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"net/http"

	"github.com/goatking91/go-gin-study/practice2/internal/app/api"
)

func Error() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()

		for _, e := range ctx.Errors {
			if _, ok := e.Err.(*validator.ValidationErrors); ok {
				api.ErrorResponse(ctx, http.StatusBadRequest, api.ErrorInvalidParams, e.Err.Error())
				return
			}
			if errors.Is(e.Err, gorm.ErrRecordNotFound) {
				api.ErrorResponse(ctx, http.StatusNotFound, api.ErrorNoExist, e.Err.Error())
				return
			}
			if _, ok := e.Err.(*mysql.MySQLError); ok {
				api.ErrorResponse(ctx, http.StatusInternalServerError, api.ErrorDb, e.Err.Error())
				return
			}
			api.ErrorResponse(ctx, http.StatusInternalServerError, api.Error, e.Err.Error())
		}
	}
}
