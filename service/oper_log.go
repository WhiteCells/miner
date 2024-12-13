package service

import (
	"errors"
	"fmt"
	"miner/dao/mysql"
	"miner/model"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type OperLogService struct {
	operLogDAO *mysql.OperLogDAO
}

func NewOperLogService() *OperLogService {
	return &OperLogService{
		operLogDAO: mysql.NewOperLogDAO(),
	}
}

// GetOperLogs 获取用户操作日志
func (s *OperLogService) GetOperLogs(ctx *gin.Context) (*[]model.OperLog, int64, error) {
	userID, exists := ctx.Value("user_id").(int)
	if !exists {
		return nil, -1, errors.New("invalid user_id in context")
	}

	// action
	action := ctx.Query("action")

	// 解析时间字符串为 time.Time 类型
	startTimeStr := ctx.Query("start_time")
	endTimeStr := ctx.Query("end_time")
	var startTime, endTime time.Time
	var err error
	if startTimeStr != "" {
		startTime, err = time.Parse(time.RFC3339, startTimeStr)
		if err != nil {
			return nil, -1, fmt.Errorf("invalid start_time format: %w", err)
		}
	}
	if endTimeStr != "" {
		endTime, err = time.Parse(time.RFC3339, endTimeStr)
		if err != nil {
			return nil, -1, fmt.Errorf("invalid end_time format: %w", err)
		}
	}

	// 分页参数解析
	pageNum, err := strconv.Atoi(ctx.Query("page_num"))
	if err != nil || pageNum <= 0 {
		return nil, -1, errors.New("invalid page_num")
	}
	pageSize, err := strconv.Atoi(ctx.Query("page_size"))
	if err != nil || pageSize <= 0 {
		return nil, -1, errors.New("invalid page_size")
	}

	query := map[string]interface{}{
		"user_id":    userID,
		"action":     action,
		"start_time": startTime,
		"end_time":   endTime,
		"page_num":   pageNum,
		"page_size":  pageSize,
	}

	// 调用 DAO 层获取日志
	logs, total, err := s.operLogDAO.GetOperLogs(query)
	if err != nil {
		return nil, -1, fmt.Errorf("failed to get oper logs: %w", err)
	}
	return logs, total, nil
}
