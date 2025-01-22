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
	route := r.Group("/fs")
	route.Use(middleware.JWTAuth())
	// route.Use(middleware.IPAuth())
	route.Use(middleware.RoleAuth(role.User, role.Admin))
	route.Use(middleware.StatusAuth())
	{
		route.POST("", fr.flightsheetController.CreateFlightsheet)
		route.DELETE("", fr.flightsheetController.DeleteFlightsheet)
		route.PUT("", fr.flightsheetController.UpdateFlightsheet)
		route.GET("", fr.flightsheetController.GetFs)
		route.PUT("/apply_wallet", fr.flightsheetController.ApplyWallet)
	}
}
