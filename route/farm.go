package route

import (
	"miner/common/role"
	"miner/controller"
	"miner/middleware"

	"github.com/gin-gonic/gin"
)

type FarmRoute struct {
	farmController *controller.FarmController
}

func NewFarmRoute() *FarmRoute {
	return &FarmRoute{
		farmController: controller.NewFarmController(),
	}
}

func (fr *FarmRoute) InitFarmRoute(r *gin.Engine) {
	route := r.Group("/farm")
	route.Use(middleware.JWTAuth())
	route.Use(middleware.IPVerify())
	route.Use(middleware.RoleAuth(role.User))
	route.Use(middleware.StatusAuth())
	{
		route.POST("", fr.farmController.CreateFarm)
		route.DELETE("", fr.farmController.DeleteFarm)
		route.PUT("", fr.farmController.UpdateFarm)
		route.GET("", fr.farmController.GetUserAllFarm)
		route.PUT("/apply_fs", fr.farmController.ApplyFlightsheet)
		route.PUT("/transfer", fr.farmController.Transfer)
	}
}
