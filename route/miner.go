package route

import (
	"miner/common/role"
	"miner/controller"
	"miner/middleware"

	"github.com/gin-gonic/gin"
)

type MinerRoute struct {
	minerController *controller.MinerController
}

func NewMinerRoute() *MinerRoute {
	return &MinerRoute{
		minerController: controller.NewMinerController(),
	}
}

func (mr *MinerRoute) InitMinerRoute(r *gin.Engine) {
	route := r.Group("/miner")
	route.Use(middleware.JWTAuth())
	route.Use(middleware.IPVerify())
	route.Use(middleware.RoleAuth(role.User))
	route.Use(middleware.StatusAuth())
	{
		route.POST("", mr.minerController.CreateMiner)
		route.DELETE("", mr.minerController.DeleteMiner)
		route.PUT("", mr.minerController.UpdateMiner)
		route.GET("", mr.minerController.GetUserAllMinerInFarm)
		route.PUT("/apply_fs", mr.minerController.ApplyFlightsheet)
		route.PUT("/transfer", mr.minerController.Transfer)
	}
}
