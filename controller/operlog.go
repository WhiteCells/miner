package controller

import (
	"miner/common/params"
	"miner/common/rsp"
	"miner/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OperlogController struct {
	operlogService *services.OperlogService
}

func NewOperLogController() *OperlogController {
	return &OperlogController{
		operlogService: services.NewOperlogService(),
	}
}

// 获取用户操作日志
func (m *OperlogController) GetOperlogs(ctx *gin.Context) {
	var params params.PageParams
	if err := ctx.ShouldBindQuery(&params); err != nil {
		rsp.Error(ctx, http.StatusBadRequest, "invalid parmas", nil)
		return
	}
	query := map[string]any{
		"page_num":  params.PageNum,
		"page_size": params.PageSize,
	}

	logs, total, err := m.operlogService.GetOperlogs(ctx, query)
	if err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	rsp.DBQuerySuccess(ctx, http.StatusOK, "get oper logs success", logs, total)
}

// 获取指定用户日志
func (m *OperlogController) GetOperlogByUserID(ctx *gin.Context) {
	var params params.PageParams
	if err := ctx.ShouldBindQuery(&params); err != nil {
		rsp.Error(ctx, http.StatusBadRequest, "invalid parmas", nil)
		return
	}
	query := map[string]any{
		"page_num":  params.PageNum,
		"page_size": params.PageSize,
	}
	userID := ctx.GetInt("user_id")

	logs, total, err := m.operlogService.GetOperlogByID(ctx, userID, query)
	if err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	rsp.DBQuerySuccess(ctx, http.StatusOK, "get oper logs success", logs, total)
}
