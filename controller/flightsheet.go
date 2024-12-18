package controller

import (
	"miner/common/dto"
	"miner/common/rsp"
	"miner/service"
	"net/http"
	"strconv"

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

// CreateFlightsheet 创建飞行表
func (c *FlightsheetController) CreateFlightsheet(ctx *gin.Context) {
	var req dto.CreateFlightsheetReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		rsp.Error(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	flightsheet, err := c.flightsheetService.CreateFlightsheet(ctx, &req)
	if err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	rsp.Success(ctx, http.StatusOK, "create flightsheet success", flightsheet)
}

// DeleteFlightsheet 删除飞行表
func (c *FlightsheetController) DeleteFlightsheet(ctx *gin.Context) {
	var req dto.DeleteFlightsheetReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		rsp.Error(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	if err := c.flightsheetService.DeleteFlightsheet(ctx, &req); err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	rsp.Success(ctx, http.StatusOK, "delete flightsheet success", nil)
}

// UpdateFlightsheet 更新飞行表
func (c *FlightsheetController) UpdateFlightsheet(ctx *gin.Context) {
	var req dto.UpdateFlightsheetReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		rsp.Error(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	if err := c.flightsheetService.UpdateFlightsheet(ctx, &req); err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	rsp.Success(ctx, http.StatusOK, "update flightsheet success", nil)
}

// GetFlightsheet 获取所有飞行表
func (c *FlightsheetController) GetFlightsheet(ctx *gin.Context) {
	pageNum, err := strconv.Atoi(ctx.Query("page_num"))
	if err != nil || pageNum <= 0 {
		rsp.Error(ctx, http.StatusBadRequest, "invalid page_numt", nil)
		return
	}
	pageSize, err := strconv.Atoi(ctx.Query("page_size"))
	if err != nil || pageSize <= 0 {
		rsp.Error(ctx, http.StatusBadRequest, "invalid page_size", nil)
		return
	}
	query := map[string]interface{}{
		"page_num":  pageNum,
		"page_size": pageSize,
	}
	flightsheets, total, err := c.flightsheetService.GetFlightsheet(ctx, query)
	if err != nil {
		rsp.Error(ctx, http.StatusInsufficientStorage, err.Error(), nil)
		return
	}

	rsp.QuerySuccess(ctx, http.StatusOK, "get user all flightsheet success", flightsheets, total)
}

// ApplyWallet 飞行表应用钱包
func (c *FlightsheetController) ApplyWallet(ctx *gin.Context) {
	var req dto.ApplyFlightsheetWalletReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		rsp.Error(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	if err := c.flightsheetService.ApplyWallet(ctx, &req); err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	rsp.Success(ctx, http.StatusOK, "apply wallet success", nil)
}
