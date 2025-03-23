package route

import (
	"miner/common/role"
	"miner/controller"
	"miner/middleware"

	"github.com/gin-gonic/gin"
)

type HiveOsRoute struct {
	hiveosController *controller.HiveosController
}

func NewHiveosRoute() *HiveOsRoute {
	return &HiveOsRoute{
		hiveosController: controller.NewHiveOsController(),
	}
}

func (m *HiveOsRoute) InitHiveosRoute(r *gin.Engine) {
	route := r.Group("")
	{
		route.POST("/worker/api", m.hiveosController.Poll)
		route.Use(middleware.JWTAuth())
		// route.Use(middleware.IPAuth())
		route.Use(middleware.RoleAuth(role.User, role.Admin))
		route.Use(middleware.StatusAuth())
		route.POST("/task", m.hiveosController.PostTask)
		route.GET("/task", m.hiveosController.GetTaskRes)
		route.GET("/stats", m.hiveosController.GetMinerStats)
		route.GET("/info", m.hiveosController.GetMinerInfo)
	}
}
