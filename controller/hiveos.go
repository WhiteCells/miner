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
	task, err := c.hiveOsService.GetTaskRes(ctx, taskID)
	if err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, "get task res failed", err.Error())
		return
	}
	rsp.Success(ctx, http.StatusOK, "get task res success", task)
}

// 获取状态
func (c *HiveOsController) GetTaskStats(ctx *gin.Context) {
	taskID := ctx.Query("task_id")
	taskStatus, err := c.hiveOsService.GetTaskStats(ctx, taskID)
	if err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, "get task status failed", err.Error())
		return
	}
	rsp.Success(ctx, http.StatusOK, "get task status success", taskStatus)
}

// 获取矿机统计信息
func (c *HiveOsController) GetMinerStats(ctx *gin.Context) {
	rigID := ctx.Query("rig_id")
	stats, err := c.hiveOsService.GetMinerStats(ctx, rigID)
	if err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, "get miner stats failed", err.Error())
		return
	}
	rsp.Success(ctx, http.StatusOK, "get miner stats success", stats)
}

// 获取矿机信息
func (c *HiveOsController) GetMinerInfo(ctx *gin.Context) {
	rigID := ctx.Query("rig_id")
	info, err := c.hiveOsService.GetMinerInfo(ctx, rigID)
	if err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, "get miner stats failed", err.Error())
		return
	}
	rsp.Success(ctx, http.StatusOK, "get miner stats success", info)
}
