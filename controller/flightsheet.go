package controller

import (
	"miner/common/dto"
	"miner/service"
	"net/http"

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
	var req dto.CreateFlightsheetReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := c.flightsheetService.CreateFlightsheet(ctx, &req); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "create flightsheet success",
	})
}

func (c *FlightsheetController) DeleteFlightsheet(ctx *gin.Context) {
	var req dto.DeleteFlightsheetReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := c.flightsheetService.DeleteFlightSheet(ctx, &req); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "delete flightsheet success",
	})
}

func (c *FlightsheetController) UpdateFlightsheet(ctx *gin.Context) {
	var req dto.UpdateFlightsheetReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := c.flightsheetService.UpdateFlightSheet(ctx, &req); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "update flightsheet success",
	})
}

func (c *FlightsheetController) GetUserAllFlightsheet(ctx *gin.Context) {
	flightsheets, err := c.flightsheetService.GetUserAllFlightsheet(ctx)
	if err != nil {
		ctx.JSON(http.StatusInsufficientStorage, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"flightsheets": flightsheets,
	})
}
