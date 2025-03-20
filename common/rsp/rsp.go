package rsp

import "github.com/gin-gonic/gin"

func LoginSuccess(ctx *gin.Context, statusCode int, msg string, data any, token string, permissions []string) {
	ctx.JSON(statusCode, gin.H{
		"code":         statusCode,
		"data":         data,
		"msg":          msg,
		"access_token": token,
		"permissions":  permissions,
	})
}

// func QuerySuccess(ctx *gin.Context, statusCode int, msg string, data any, total int64) {
func QuerySuccess(ctx *gin.Context, statusCode int, msg string, data any) {
	ctx.JSON(statusCode, gin.H{
		"code": statusCode,
		"data": data,
		"msg":  msg,
	})
}

func DBQuerySuccess(ctx *gin.Context, statusCode int, msg string, data any, total int64) {
	ctx.JSON(statusCode, gin.H{
		"code":  statusCode,
		"data":  data,
		"total": total,
		"msg":   msg,
	})
}

func Success(ctx *gin.Context, statusCode int, msg string, data any) {
	ctx.JSON(statusCode, gin.H{
		"code": statusCode,
		"data": data,
		"msg":  msg,
	})
}

func Error(ctx *gin.Context, statusCode int, msg string, data any) {
	ctx.JSON(statusCode, gin.H{
		"code": statusCode,
		"data": data,
		"msg":  msg,
	})
}
