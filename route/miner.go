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
	route.Use(middleware.IPVerify())
	route.Use(middleware.JWTAuth())
	route.Use(middleware.RoleAuth(role.User))
	{
		route.POST("", mr.minerController.CreateMiner)
		route.DELETE("", mr.minerController.DeleteMiner)
		route.PUT("", mr.minerController.UpdateMiner)
		route.GET("", mr.minerController.GetUserAllMinerInFarm)
		route.POST("/apply_fs", mr.minerController.ApplyFlightsheet)
	}
}
