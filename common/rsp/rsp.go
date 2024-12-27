package rsp

import "github.com/gin-gonic/gin"

func LoginSuccess(ctx *gin.Context, statusCode int, msg string, data interface{}, token string) {
	ctx.JSON(statusCode, gin.H{
		"code":         statusCode,
		"data":         data,
		"msg":          msg,
		"access_token": token,
	})
}

// func QuerySuccess(ctx *gin.Context, statusCode int, msg string, data interface{}, total int64) {
func QuerySuccess(ctx *gin.Context, statusCode int, msg string, data interface{}) {
	ctx.JSON(statusCode, gin.H{
		"code": statusCode,
		"data": data,
		"msg":  msg,
	})
}

func DBQuerySuccess(ctx *gin.Context, statusCode int, msg string, data interface{}, total int64) {
	ctx.JSON(statusCode, gin.H{
		"code":  statusCode,
		"data":  data,
		"total": total,
		"msg":   msg,
	})
}

func Success(ctx *gin.Context, statusCode int, msg string, data interface{}) {
	ctx.JSON(statusCode, gin.H{
		"code": statusCode,
		"data": data,
		"msg":  msg,
	})
}

func Error(ctx *gin.Context, statusCode int, msg string, data interface{}) {
	ctx.JSON(statusCode, gin.H{
		"code": statusCode,
		"data": data,
		"msg":  msg,
	})
}
