package controller

import (
	"miner/service"

	"github.com/gin-gonic/gin"
)

type FlightsheetController struct {
	flightsheetService *service.FlightsheetService
}

func NewFlightsheetController() *FlightsheetController {
	return &FlightsheetController{
		flightsheetService: service.NewFlightsheetService(),
	}
}

func (c *FlightsheetController) CreateFlightsheet(ctx *gin.Context) {
	// var req dto.
}

func (c *FlightsheetController) DeleteFlightsheet(ctx *gin.Context) {
	//
}

func (c *FlightsheetController) UpdateFlightsheet(ctx *gin.Context) {
	//
}

func (c *FlightsheetController) GetUserAllFlightsheet(ctx *gin.Context) {
	//
}
