package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"strings"
	"time"

	"github.com/goatking91/go-gin-study/practice2/internal/app/api"
	"github.com/goatking91/go-gin-study/practice2/pkg/logger"
)

// HttpLogger access log 를 출력
// [INFO]2022/05/04 16:18:28 [>> a463ed80-1cd5-4a95-9c20-24b36c27f456|400|0400|::1|6.620958ms|236|POST|"/api/v1/u/crypto/hash/sha512/aes256"]
func HttpLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		txid := uuid.New().String()

		baseResponse := &api.BaseRes{
			TxID:   txid,
			Path:   c.FullPath(),
			Method: c.Request.Method,
		}

		c.Set("BASE_RESPONSE", baseResponse)

		// before request
		c.Next()

		// after request
		// do somthing

		// access the status we are sending
		status := c.Writer.Status()
		size := c.Writer.Size()
		latency := time.Since(t)
		code := c.Writer.Header().Get("X-CODE")
		//txid := c.Writer.Header().Get("X-TXID")

		// healthcheck 인 경우 로깅에서 제외
		if !strings.HasPrefix(c.FullPath(), "/v1/ping") {
			logger.S.Infof(">> %s|%d|%s|%v|%s|%d|%s|\"%s\"",
				txid, status, code, c.ClientIP(), latency, size, c.Request.Method, c.Request.RequestURI)
		}
	}
}
