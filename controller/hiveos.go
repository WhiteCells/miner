package controller

import (
	"miner/common/dto"
	"miner/common/rsp"
	"miner/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HiveosController struct {
	hiveosService *services.HiveosService
}

func NewHiveOsController() *HiveosController {
	return &HiveosController{
		hiveosService: services.NewHiveosService(),
	}
}

// 轮询
func (m *HiveosController) Poll(ctx *gin.Context) {
	m.hiveosService.Poll(ctx)
}

// 发送任务
func (c *HiveosController) PostTask(ctx *gin.Context) {
	var req dto.PostTaskReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		rsp.Error(ctx, http.StatusBadRequest, "req fromat failed", err.Error())
		return
	}
	taskID, err := c.hiveosService.PostTask(ctx, &req)
	if err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, "post task failed", err.Error())
		return
	}
	rsp.Success(ctx, http.StatusOK, "post task success", taskID)
}

// 获取命令结果
func (c *HiveosController) GetTaskRes(ctx *gin.Context) {
	taskID := ctx.Query("task_id")
	task, err := c.hiveosService.GetTaskRes(ctx, taskID)
	if err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, "get task res failed", err.Error())
		return
	}
	rsp.Success(ctx, http.StatusOK, "get task res success", task)
}

// 获取状态
func (c *HiveosController) GetTaskStats(ctx *gin.Context) {
	taskID := ctx.Query("task_id")
	taskStatus, err := c.hiveosService.GetTaskStats(ctx, taskID)
	if err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, "get task status failed", err.Error())
		return
	}
	rsp.Success(ctx, http.StatusOK, "get task status success", taskStatus)
}

// 获取矿机统计信息
func (c *HiveosController) GetMinerStats(ctx *gin.Context) {
	rigID := ctx.Query("rig_id")
	stats, err := c.hiveosService.GetMinerStats(ctx, rigID)
	if err != nil {
		if err.Error() == "redis: nil" {
			rsp.Success(ctx, http.StatusOK, "get miner stats success", nil)
			return
		}
		rsp.Error(ctx, http.StatusInternalServerError, "get miner stats failed", err.Error())
		return
	}
	rsp.Success(ctx, http.StatusOK, "get miner stats success", stats)
}

// 获取矿机信息
func (c *HiveosController) GetMinerInfo(ctx *gin.Context) {
	rigID := ctx.Query("rig_id")
	info, err := c.hiveosService.GetMinerInfo(ctx, rigID)
	if err != nil {
		if err.Error() == "redis: nil" {
			rsp.Success(ctx, http.StatusOK, "get miner stats success", nil)
			return
		}
		rsp.Error(ctx, http.StatusInternalServerError, "get miner info failed", err.Error())
		return
	}
	rsp.Success(ctx, http.StatusOK, "get miner stats success", info)
}
