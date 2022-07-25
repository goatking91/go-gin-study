package api

import "github.com/gin-gonic/gin"

type BaseRes struct {
	TxID   string `json:"txid"`
	Path   string `json:"path"`
	Method string `json:"method"`
}

type SuccessRes struct {
	*BaseRes
	Data interface{} `json:"data"`
}

type Err struct {
	Code        string `json:"code"`
	Message     string `json:"message"`
	Description string `json:"description"`
}

type ErrorRes struct {
	*BaseRes
	Error Err `json:"error"`
}

func Response(ctx *gin.Context, status int, Data interface{}) {
	baseResponse, _ := ctx.Get("BASE_RESPONSE")
	response := SuccessRes{
		BaseRes: baseResponse.(*BaseRes),
		Data:    Data,
	}
	ctx.IndentedJSON(status, response)
}

func ErrorResponse(ctx *gin.Context, status int, errorCode string, description string) {
	baseResponse, _ := ctx.Get("BASE_RESPONSE")
	response := ErrorRes{
		BaseRes: baseResponse.(*BaseRes),
		Error: Err{
			Code:        errorCode,
			Message:     Messages[errorCode],
			Description: description,
		},
	}
	ctx.AbortWithStatusJSON(status, response)
}
