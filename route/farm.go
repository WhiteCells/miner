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

func (fr *FarmRoute) InitFarmRoute(rg *gin.Engine) {
	route := rg.Group("/farm")
	route.Use(middleware.IPVerify())
	route.Use(middleware.JWTAuth())
	route.Use(middleware.OperLogger())
	route.Use(middleware.RoleAuth(role.User))
	{
		route.POST("/add", fr.farmController.CreateFarm)
	}
}
