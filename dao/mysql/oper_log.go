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

// 创建操作日志
func (dao *OperLogDAO) CreateOperLog(log *model.OperLog) error {
	return utils.DB.Create(log).Error
}

// 获取操作日志列表
func (dao *OperLogDAO) GetOperLogs(query map[string]interface{}, page, pageSize int) ([]model.OperLog, int64, error) {
	var logs []model.OperLog
	var total int64

	db := utils.DB.Model(&model.OperLog{})

	// 添加查询条件
	if userID, ok := query["user_id"].(int); ok {
		db = db.Where("user_id = ?", userID)
	}
	if action, ok := query["action"].(string); ok {
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
		return nil, 0, err
	}

	// 获取分页数据
	err := db.Offset((page - 1) * pageSize).
		Limit(pageSize).
		Order("created_at DESC").
		Find(&logs).Error

	return logs, total, err
}
