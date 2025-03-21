package middleware

import (
	"bytes"
	"encoding/json"
	"io"
	"time"

	"miner/model"
	"miner/utils"

	"github.com/gin-gonic/gin"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

// 操作日志中间件
func OperLog() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// GET请求不记录日志
		if ctx.Request.Method == "GET" {
			ctx.Next()
			return
		}

		// 获取请求体
		var requestBody []byte
		if ctx.Request.Body != nil {
			requestBody, _ = io.ReadAll(ctx.Request.Body)
			ctx.Request.Body = io.NopCloser(bytes.NewBuffer(requestBody))
		}

		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: ctx.Writer}
		ctx.Writer = blw

		// 继续处理业务
		ctx.Next()

		// 业务处理完后
		userID, exists := ctx.Value("user_id").(int)
		if !exists {
			// ctx.
			return
		}

		// 创建操作日志
		operLog := model.Operlog{
			UserID: userID,
			Time:   time.Now(),
			Action: ctx.Request.Method,
			Target: ctx.FullPath(),
			IP:     ctx.ClientIP(),
			Agent:  ctx.Request.UserAgent(),
			Status: ctx.Writer.Status(),
		}

		// 记录请求详情
		detail := map[string]any{
			"request":  string(requestBody),
			"response": blw.body.String(),
		}
		detailJSON, _ := json.Marshal(detail)
		operLog.Detail = string(detailJSON)

		// 异步保存日志
		go func() {
			if err := utils.DB.Create(&operLog).Error; err != nil {
				utils.Logger.Error("Failed to save operation log, " + err.Error())
			}
		}()

		ctx.Next()
	}
}

// 登陆日志中间件
func LoginLog() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()

		userID, exists := ctx.Value("user_id").(int)
		if !exists {
			return
		}

		loginLog := &model.Loginlog{
			UserID: userID,
			Time:   time.Now(),
			IP:     ctx.ClientIP(),
			Status: ctx.Writer.Status(),
		}

		go func() {
			if err := utils.DB.Create(&loginLog).Error; err != nil {
				utils.Logger.Error("Failed to save login log, " + err.Error())
			}
		}()

		ctx.Next()
	}
}
