package mysql

import (
	"miner/model"
	"miner/utils"
	"time"
)

type OperLogDAO struct{}

func NewOperLogDAO() *OperLogDAO {
	return &OperLogDAO{}
}

// GetOperLogs 获取操作日志
func (dao *OperLogDAO) GetOperLogs(query map[string]interface{}) (*[]model.OperLog, int64, error) {
	var logs []model.OperLog
	var total int64

	db := utils.DB.Model(&model.OperLog{})

	// 添加查询条件
	if userID, ok := query["user_id"].(int); ok {
		db = db.Where("user_id = ?", userID)
	}
	if action, ok := query["action"].(string); ok && action != "" {
		db = db.Where("action = ?", action)
	}
	if startTime, ok := query["start_time"].(time.Time); ok {
		db = db.Where("time >= ?", startTime)
	}
	if endTime, ok := query["end_time"].(time.Time); ok {
		db = db.Where("time <= ?", endTime)
	}

	// 获取总数
	if err := db.Count(&total).Error; err != nil {
		return nil, -1, err
	}

	pageNum := query["page_num"].(int)
	pageSize := query["page_size"].(int)

	// 获取分页数据
	err := db.Offset((pageNum - 1) * pageSize).
		Limit(pageSize).
		Order("time").
		Find(&logs).Error

	return &logs, total, err
}
