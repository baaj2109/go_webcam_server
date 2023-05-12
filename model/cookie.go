package model

import (
	"time"

	"github.com/gin-gonic/gin"
)

type cookie struct{}

var Cookie = &cookie{}

func (c cookie) Set(ctx *gin.Context, token string) {
	cookie, err := ctx.Cookie("user")
	if err != nil {
		ctx.SetCookie("user", token, time.Now().Hour()*3, "/", "localhost", false, false)
		ctx.JSON(200, gin.H{"msg": "set cookie successfully"})
	} else {
		ctx.JSON(200, gin.H{"msg": cookie})
	}
}

func (c cookie) Get(ctx *gin.Context) string {
	cookie, err := ctx.Cookie("user")
	if err != nil {
		return ""
	}
	return cookie
}

func (c cookie) Del(ctx *gin.Context) {
	ctx.SetCookie("user", "", -1, "/", "localhost", false, false)
	ctx.JSON(200, gin.H{"msg": "delete cookie successfully"})
}
