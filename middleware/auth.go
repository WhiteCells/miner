package middleware

import (
	"context"
	"net/http"
	"strings"

	"miner/common/role"
	"miner/common/rsp"
	"miner/common/status"
	"miner/dao/redis"
	"miner/utils"

	"github.com/gin-gonic/gin"
)

var UserRDB = redis.NewUserRDB()

// JWT 验证
func JWTAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			rsp.Error(ctx, http.StatusUnauthorized, "authorization header is required", nil)
			ctx.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			rsp.Error(ctx, http.StatusUnauthorized, "authorization header format error", nil)
			ctx.Abort()
			return
		}

		// token 是否存在 ban tokne 中
		if UserRDB.ExistsBanToken(ctx, parts[1]) {
			rsp.Error(ctx, http.StatusUnauthorized, "token expired", nil)
			ctx.Abort()
			return
		}

		claims, err := utils.ParseToken(parts[1])
		if err != nil {
			rsp.Error(ctx, http.StatusUnauthorized, err.Error(), nil)
			ctx.Abort()
			return
		}

		// 保存用户信息到 ctx
		ctx.Set("user_id", claims.UserID)
		ctx.Set("user_name", claims.Username)

		ctx.Next()
	}
}

// 角色验证
func RoleAuth(roles ...role.RoleType) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userID, exists := ctx.Value("user_id").(string)
		if !exists {
			rsp.Error(ctx, http.StatusForbidden, "user_id not found", nil)
			ctx.Abort()
			return
		}
		user, err := UserRDB.GetByID(ctx, userID)
		if err != nil {
			rsp.Error(ctx, http.StatusForbidden, "user not found", nil)
			ctx.Abort()
			return
		}
		for _, role := range roles {
			if role == user.Role {
				ctx.Next()
				return
			}
		}
		rsp.Error(ctx, http.StatusForbidden, "permission denied", nil)
		ctx.Abort()
	}
}

// 状态验证
func StatusAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userID, exists := ctx.Value("user_id").(string)
		if !exists {
			rsp.Error(ctx, http.StatusForbidden, "user_id not in context", nil)
			ctx.Abort()
			return
		}
		user, err := UserRDB.GetByID(ctx, userID)
		if err != nil {
			rsp.Error(ctx, http.StatusForbidden, "user not found", nil)
			ctx.Abort()
			return
		}

		if user.Status != status.UserOn {
			rsp.Error(ctx, http.StatusBadRequest, "user status off", nil)
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}

var SystemRDB = redis.NewAdminRDB()

// 注册开关验证
func RegisterAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// b, err := SystemRDB.GetSwitchRegister(ctx)
		// if err != nil {
		// 	rsp.Error(ctx, http.StatusForbidden, "system rdb not found", nil)
		// 	ctx.Abort()
		// 	return
		// }
		// if b != string(status.RegisterOn) {
		// 	rsp.Error(ctx, http.StatusBadRequest, "register closed", nil)
		// 	ctx.Abort()
		// 	return
		// }
		ctx.Next()
	}
}

// IPAuth 验证
func IPAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userID, exists := ctx.Value("user_id").(string)
		if !exists {
			rsp.Error(ctx, http.StatusUnauthorized, "unauthorized", nil)
			ctx.Abort()
			return
		}

		// 从 Redis 获取上次登录 IP
		c := context.Background()
		user, err := UserRDB.GetByID(c, userID)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "User info not found",
			})
			ctx.Abort()
			return
		}

		// TODO 如果 IP 不同，需要重新验证
		if user.LastLoginIP != ctx.ClientIP() {
			rsp.Error(ctx, http.StatusUnauthorized, "New IP detected", nil)
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
