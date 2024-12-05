package middleware

import (
	"context"
	"fmt"
	"miner/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// IPVerify IP 验证中间件
func IPVerify() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, exists := c.Get("user_id")
		if !exists {
			return
		}

		currentIP := c.ClientIP()
		ctx := context.Background()

		// 从 Redis 获取上次登录 IP
		key := fmt.Sprintf("user:%d:last_ip", userID.(int))
		lastIP, err := utils.RDB.Get(ctx, key)
		if err != nil {
			// 首次登录，记录IP
			utils.RDB.Set(ctx, key, currentIP, 24*time.Hour)
			return
		}

		// 如果 IP 不同，需要重新验证 todo
		if lastIP != currentIP {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "New IP detected",
			})
			c.Abort()
			return
		}
	}
}
