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

func GetOperLogSuccess(ctx *gin.Context, statusCode int, msg string, data interface{}, total int64) {
	ctx.JSON(statusCode, gin.H{
		"code":  statusCode,
		"total": total,
		"data":  data,
		"msg":   msg,
	})
}

func GetPointsRecordsSuccess(ctx *gin.Context, statusCode int, msg string, data interface{}, total int64) {
	ctx.JSON(statusCode, gin.H{
		"code":  statusCode,
		"total": total,
		"data":  data,
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
