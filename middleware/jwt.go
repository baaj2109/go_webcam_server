package middleware

import (
	"net/http"
	"time"

	"github.com/baaj2109/webcam_server/config"
	"github.com/baaj2109/webcam_server/utils"

	"github.com/gin-gonic/gin"
)

func JWT(cfg *config.JWTConfig) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		token := ctx.Query("token")
		if token == "" {
			ctx.JSON(http.StatusOK, gin.H{
				"status": -1,
				"msg":    "failed to get token",
				"data":   nil,
			})
			ctx.Abort()
			return
		} else {
			// 解析token
			claims, err := utils.ParseToken(token, cfg)
			if err != nil {
				ctx.JSON(http.StatusOK, gin.H{
					"status": -1,
					"msg":    "ERROR_AUTH_CHECK_TOKEN_FAIL",
					"data":   nil,
				})
				ctx.Abort()
				return
			} else if time.Now().Unix() > claims.ExpiresAt {
				ctx.JSON(http.StatusOK, gin.H{
					"status": -1,
					"msg":    "ERROR_AUTH_CHECK_TOKEN_TIMEOUT",
					"data":   nil,
				})
				ctx.Abort()
				return
			}
		}
		ctx.Next()
	}
}
