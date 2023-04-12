package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetHome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "current time: " + time.Now().GoString(),
	})
}
