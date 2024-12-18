package middleware

import (
	"miner/common/rsp"
	"miner/model"
	"miner/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// var userCache = redis.NewUserCache()

// IPVerify IP 验证中间件
func IPVerify() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userID, exists := ctx.Value("user_id").(int)
		if !exists {
			rsp.Error(ctx, http.StatusUnauthorized, "unauthorized", nil)
			ctx.Abort()
			return
		}

		currentIP := ctx.ClientIP()

		// 从 Redis 获取上次登录 IP
		// c := context.Background()
		// user, err := userCache.GetUserInfoByID(c, userID)
		// if err != nil {
		// 	ctx.JSON(http.StatusUnauthorized, gin.H{
		// 		"code": 401,
		// 		"msg":  "User info not found",
		// 	})
		// 	ctx.Abort()
		// 	return
		// }

		var user model.User
		if err := utils.DB.Where("id = ?", userID).First(&user).Error; err != nil {
			rsp.Error(ctx, http.StatusInternalServerError, "user not found", nil)
			ctx.Abort()
			return
		}

		// TODO 如果 IP 不同，需要重新验证
		if user.LastLoginIP != currentIP {
			rsp.Error(ctx, http.StatusUnauthorized, "New IP detected", nil)
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
