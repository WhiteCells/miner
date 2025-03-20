package mysql

import (
	"miner/model"
	"miner/utils"
	"strconv"
)

type OperLogDAO struct{}

func NewOperLogDAO() *OperLogDAO {
	return &OperLogDAO{}
}

// GetOperLogs 获取操作日志
func (dao *OperLogDAO) GetOperLogs(query map[string]any) (*[]model.Operlog, int64, error) {
	var logs []model.Operlog
	var total int64

	db := utils.DB.Model(&model.Operlog{})

	// 添加查询条件
	if userIDStr, ok := query["user_id"].(string); ok {
		if userID, err := strconv.Atoi(userIDStr); err == nil {
			db = db.Where("user_id = ?", userID)
		} else {
			return nil, -1, err
		}
	}
	// if action, ok := query["action"].(string); ok && action != "" {
	// 	db = db.Where("action = ?", action)
	// }
	// if startTime, ok := query["start_time"].(time.Time); ok {
	// 	db = db.Where("time >= ?", startTime)
	// }
	// if endTime, ok := query["end_time"].(time.Time); ok {
	// 	db = db.Where("time <= ?", endTime)
	// }

	// 获取总数
	if err := db.Count(&total).Error; err != nil {
		return nil, -1, err
	}

	pageNum := query["page_num"].(int)
	pageSize := query["page_size"].(int)

	// 获取分页数据
	err := db.Offset((pageNum - 1) * pageSize).
		Limit(pageSize).
		Order("time desc").
		Find(&logs).Error

	return &logs, total, err
}
