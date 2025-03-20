package controller

import (
	"miner/common/rsp"
	"miner/service"
	"net/http"
	"strconv"

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
	userID, exists := ctx.Value("user_id").(string)
	if !exists {
		rsp.Error(ctx, http.StatusBadRequest, "invalid user_id in context", nil)
		return
	}

	// action
	// action := ctx.Query("action")

	// 解析时间字符串为 time.Time 类型
	// startTimeStr := ctx.Query("start_time")
	// endTimeStr := ctx.Query("end_time")
	// var startTime, endTime time.Time
	// var err error
	// if startTimeStr != "" {
	// 	startTime, err = time.Parse(time.RFC3339, startTimeStr)
	// 	if err != nil {
	// 		rsp.Error(ctx, http.StatusBadRequest, "invalid start_time format", nil)
	// 		return
	// 	}
	// }
	// if endTimeStr != "" {
	// 	endTime, err = time.Parse(time.RFC3339, endTimeStr)
	// 	if err != nil {
	// 		rsp.Error(ctx, http.StatusBadRequest, "invalid end_time format", nil)
	// 		return
	// 	}
	// }

	// 分页参数解析
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

	query := map[string]any{
		"user_id": userID,
		// "action":  action,
		// "start_time": startTime,
		// "end_time":   endTime,
		"page_num":  pageNum,
		"page_size": pageSize,
	}

	logs, total, err := c.operLogService.GetOperLogs(ctx, query)
	if err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	rsp.DBQuerySuccess(ctx, http.StatusOK, "get oper logs success", logs, total)
}
