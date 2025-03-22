package mysql

import (
	"context"
	"miner/model"
	"miner/utils"
)

type LoginlogDAO struct {
}

func NewLoginlogDAO() *LoginlogDAO {
	return &LoginlogDAO{}
}

// 获取指定用户日志
func (m *LoginlogDAO) GetLoginlogByID(ctx context.Context, userID int, query map[string]any) (*[]model.Loginlog, int64, error) {
	var logs []model.Loginlog
	var total int64

	db := utils.DB.WithContext(ctx).
		Model(&model.Loginlog{}).
		Where("user_id=?", userID)

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

// 获取登录日志
func (m *LoginlogDAO) GetLoginlogs(ctx context.Context, query map[string]any) (*[]model.Loginlog, int64, error) {
	var logs []model.Loginlog
	var total int64

	db := utils.DB.WithContext(ctx).
		Model(&model.Loginlog{})

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
