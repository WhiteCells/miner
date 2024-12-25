package route

import (
	"miner/common/role"
	"miner/middleware"

	"github.com/gin-gonic/gin"
)

type HiveosRoute struct {
}

func NewHiveosRoute() *HiveosRoute {
	return &HiveosRoute{}
}

func (hr *HiveosRoute) InitHiveosRoute(r *gin.Engine) {
	route := r.Group("/hiveos")
	route.Use(middleware.JWTAuth())
	route.Use(middleware.IPVerify())
	route.Use(middleware.RoleAuth(role.User))
	route.Use(middleware.StatusAuth())
	{
		route.POST("/worker/api", func(ctx *gin.Context) {

		})
	}
}
