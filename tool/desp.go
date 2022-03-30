package tool

import "github.com/gin-gonic/gin"

func Success(ctx *gin.Context, code int, info string) {
	ctx.JSON(code, gin.H{
		"code": code,
		"info": info,
	})
}

func Failure(ctx *gin.Context, code int, info string) {
	ctx.JSON(code, gin.H{
		"code": code,
		"info": info,
	})
}
