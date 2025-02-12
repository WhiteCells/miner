package route

import (
	"miner/common/role"
	"miner/controller"
	"miner/middleware"

	"github.com/gin-gonic/gin"
)

type HiveOsRoute struct {
	hiveOsController *controller.HiveOsController
}

func NewHiveosRoute() *HiveOsRoute {
	return &HiveOsRoute{
		hiveOsController: controller.NewHiveOsController(),
	}
}

func (hr *HiveOsRoute) InitHiveosRoute(r *gin.Engine) {
	route := r.Group("")
	{
		route.POST("/worker/api", hr.hiveOsController.Poll)
		route.Use(middleware.JWTAuth())
		// route.Use(middleware.IPAuth())
		route.Use(middleware.RoleAuth(role.User, role.Admin))
		route.Use(middleware.StatusAuth())
		route.POST("/task", hr.hiveOsController.PostTask)
		route.GET("/task", hr.hiveOsController.GetTaskRes)
		route.GET("/stats", hr.hiveOsController.GetMinerStats)
		route.GET("/info", hr.hiveOsController.GetMinerInfo)
	}
}
