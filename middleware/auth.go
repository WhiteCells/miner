package middleware

import (
	"net/http"
	"strings"

	"miner/common/role"
	"miner/common/rsp"
	"miner/common/status"
	"miner/model"
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

		ctx.Next()
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
				ctx.Next()
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

// 状态验证
func StatusAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userID, exists := ctx.Value("user_id").(int)
		if !exists {
			rsp.Error(ctx, http.StatusForbidden, "user_id not in context", nil)
			ctx.Abort()
			return
		}
		// 通过 userID 查找用户状态
		// 先从缓存中查找，缓存未命中再从数据库中查找

		var user model.User
		if err := utils.DB.Where("id = ?", userID).Find(&user).Error; err != nil {
			rsp.Error(ctx, http.StatusInternalServerError, "user not found", nil)
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
