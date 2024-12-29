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
	route := r.Group("/hiveos")
	route.Use(middleware.JWTAuth())
	route.Use(middleware.IPAuth())
	route.Use(middleware.RoleAuth(role.User, role.Admin))
	route.Use(middleware.StatusAuth())
	{
		route.POST("/worker/api", hr.hiveOsController.Interact)
		route.POST("/cmd", hr.hiveOsController.SendCmd)
		route.GET("/cmd", hr.hiveOsController.GetCmdRes)
		route.POST("/config", hr.hiveOsController.SetConfig)
		route.GET("/config", hr.hiveOsController.GetConfigRes)
	}
}
