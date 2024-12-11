package rsp

import "github.com/gin-gonic/gin"

func SuccessRsp(ctx *gin.Context, statusCode int, msg string, data interface{}) {
	ctx.JSON(statusCode, gin.H{
		"code": statusCode,
		"data": data,
		"msg":  msg,
	})
}

func ErrorRsp(ctx *gin.Context, statusCode int, msg string, data interface{}) {
	ctx.JSON(statusCode, gin.H{
		"code": statusCode,
		"data": data,
		"msg":  msg,
	})
}
