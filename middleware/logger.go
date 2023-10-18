package middleware

import (
	"time"

	"github.com/baaj2109/webcam_server/utils"
	"github.com/gin-gonic/gin"
)

func LoggerToFile() gin.HandlerFunc {
	logger := utils.Logger()
	return func(c *gin.Context) {
		startTime := time.Now()               // 開始時間
		endTime := time.Now()                 // 结束時間
		latencyTime := endTime.Sub(startTime) // 執行時間
		reqMethod := c.Request.Method         // 請求方式
		reqUri := c.Request.RequestURI        // 請求路由
		statusCode := c.Writer.Status()       // 狀態
		clientIP := c.ClientIP()              // 請求IP

		//日誌格式
		logger.Infof("| %3d | %13v | %15s | %s | %s |",
			statusCode,
			latencyTime,
			clientIP,
			reqMethod,
			reqUri,
		)
		c.Next() // 處理請求
	}
}
