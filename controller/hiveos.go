package controller

import (
	"miner/common/dto"
	"miner/common/rsp"
	"miner/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HiveOsController struct {
	hiveOsService *service.HiveOsService
}

func NewHiveOsController() *HiveOsController {
	return &HiveOsController{
		hiveOsService: service.NewHiveOsService(),
	}
}

// 轮询
func (c *HiveOsController) Poll(ctx *gin.Context) {
	c.hiveOsService.Poll(ctx)
}

// 发送任务
func (c *HiveOsController) PostTask(ctx *gin.Context) {
	var req dto.PostTaskReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		rsp.Error(ctx, http.StatusBadRequest, "req fromat failed", err.Error())
		return
	}
	taskID, err := c.hiveOsService.PostTask(ctx, &req)
	if err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, "post task failed", err.Error())
		return
	}
	rsp.Success(ctx, http.StatusOK, "post task success", taskID)
}

// 获取命令结果
func (c *HiveOsController) GetTaskRes(ctx *gin.Context) {
	taskID := ctx.Query("task_id")
	res, err := c.hiveOsService.GetTaskRes(ctx, taskID)
	if err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, "req format failed", err.Error())
		return
	}
	rsp.Success(ctx, http.StatusOK, "get task res success", res)
}

// 获取状态
func (c *HiveOsController) GetStats(ctx *gin.Context) {

}
