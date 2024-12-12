package route

import (
	"miner/common/role"
	"miner/controller"
	"miner/middleware"

	"github.com/gin-gonic/gin"
)

type FlightsheetRoute struct {
	flightsheetController *controller.FlightsheetController
}

func NewFlightsheetRoute() *FlightsheetRoute {
	return &FlightsheetRoute{
		flightsheetController: controller.NewFlightsheetController(),
	}
}

func (fr *FlightsheetRoute) InitFlightsheetRoute(r *gin.Engine) {
	route := r.Group("/flightsheet")
	route.Use(middleware.IPVerify())
	route.Use(middleware.JWTAuth())
	route.Use(middleware.RoleAuth(role.User))
	{
		route.POST("", fr.flightsheetController.CreateFlightsheet)
		route.DELETE("", fr.flightsheetController.DeleteFlightsheet)
		route.PUT("", fr.flightsheetController.UpdateFlightsheet)
		route.GET("", fr.flightsheetController.GetUserAllFlightsheet)
		route.POST("/apply_wallet", fr.flightsheetController.ApplyWallet)
	}
}
