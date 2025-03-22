package mysql

import (
	"context"
	"miner/model"
	"miner/utils"
)

type OperLogDAO struct{}

func NewOperLogDAO() *OperLogDAO {
	return &OperLogDAO{}
}

// 获取操作日志
func (OperLogDAO) GetOperlogs(ctx context.Context, query map[string]any) (*[]model.Operlog, int64, error) {
	var logs []model.Operlog
	var total int64

	pageNum := query["page_num"].(int)
	pageSize := query["page_size"].(int)

	db := utils.DB.WithContext(ctx).
		Model(&model.Operlog{})

	if err := db.Count(&total).Error; err != nil {
		return nil, -1, err
	}

	err := db.Offset((pageNum - 1) * pageSize).
		Limit(pageSize).
		Order("time desc").
		Find(&logs).Error

	return &logs, total, err
}

// 获取指定用户操作日志
func (dao *OperLogDAO) GetOperlogByID(ctx context.Context, userID int, query map[string]any) (*[]model.Operlog, int64, error) {
	var logs []model.Operlog
	var total int64

	pageNum := query["page_num"].(int)
	pageSize := query["page_size"].(int)

	db := utils.DB.WithContext(ctx).
		Model(&model.Operlog{}).
		Where("user_id=?", userID)

	if err := db.Count(&total).Error; err != nil {
		return nil, -1, err
	}

	err := db.Offset((pageNum - 1) * pageSize).
		Limit(pageSize).
		Order("time desc").
		Find(&logs).Error

	return &logs, total, err
}
