package mysql

import (
	"miner/model"
	"miner/utils"
)

type LoginlogDAO struct {
}

func NewLoginlogDAO() *LoginlogDAO {
	return &LoginlogDAO{}
}

// 获取登录日志
func (m *LoginlogDAO) GetLoginlogs(query map[string]any) (*[]model.Loginlog, int64, error) {
	var logs []model.Loginlog
	var total int64

	db := utils.DB.Where(&model.Loginlog{})

	if userID, ok := query["user_id"].(int); ok {
		db = db.Where("user_id=?", userID)
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, -1, err
	}

	pageNum := query["page_num"].(int)
	pageSize := query["page_size"].(int)

	err := db.Offset((pageNum - 1) * pageSize).
		Limit(pageSize).
		Find(&logs).Error

	return &logs, total, err
}
