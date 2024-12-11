package middleware

import (
	"net/http"
	"strings"

	"miner/common/role"
	"miner/utils"

	"github.com/gin-gonic/gin"
)

// JWT 认证
func JWTAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "Authorization header is required",
			})
			ctx.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "Authorization header format error",
			})
			ctx.Abort()
			return
		}

		claims, err := utils.ParseToken(parts[1])
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			ctx.Abort()
			return
		}

		// 保存用户信息到 ctx
		ctx.Set("user_id", claims.UserID)
		ctx.Set("user_name", claims.Username)
		ctx.Set("user_role", claims.Role)
	}
}

// 角色验证
func RoleAuth(roles ...role.RoleType) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctxRole, exists := ctx.Get("user_role")
		if !exists || ctxRole == "" {
			ctx.JSON(http.StatusForbidden, gin.H{
				"code": 403,
				"msg":  "Role information not found",
			})
			ctx.Abort()
			return
		}

		for _, role := range roles {
			if role == ctxRole {
				return
			}
		}

		ctx.JSON(http.StatusForbidden, gin.H{
			"code": 403,
			"msg":  "Permission denied",
		})
		ctx.Abort()
	}
}
