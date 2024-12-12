package middleware

import (
	"context"
	"miner/dao/redis"
	"net/http"

	"github.com/gin-gonic/gin"
)

var userCache = redis.NewUserCache()

// IPVerify IP 验证中间件
func IPVerify() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, exists := c.Value("user_id").(int)
		if !exists {
			// c.JSON(http.StatusUnauthorized, gin.H{
			// 	"code": 401,
			// 	"msg":  "Unauthorized",
			// })
			// c.Abort()
			return
		}

		currentIP := c.ClientIP()
		ctx := context.Background()

		// 从 Redis 获取上次登录 IP
		user, err := userCache.GetUserInfoByID(ctx, userID)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "User info not found",
			})
			c.Abort()
			return
		}

		// TODO 如果 IP 不同，需要重新验证
		if user.LastLoginIP != currentIP {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "New IP detected",
			})
			c.Abort()
			return
		}
	}
}
