package controller

import (
	"miner/common/rsp"
	"miner/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OperLogController struct {
	operLogService *service.OperLogService
}

func NewOperLogController() *OperLogController {
	return &OperLogController{
		operLogService: service.NewOperLogService(),
	}
}

// GetOperLogs 获取用户日志
func (c *OperLogController) GetOperLogs(ctx *gin.Context) {
	// var req dto.GetUserOperLogsReq
	// if err := ctx.ShouldBindJSON(&req); err != nil {
	// 	rsp.Error(ctx, http.StatusBadRequest, err.Error(), nil)
	// 	return
	// }

	logs, total, err := c.operLogService.GetOperLogs(ctx)
	if err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	rsp.GetOperLogSuccess(ctx, http.StatusOK, "get oper logs success", logs, total)
}
