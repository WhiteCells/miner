package mysql

import (
	"context"
	"miner/model"
	"miner/utils"
)

type PointslogDAO struct{}

func NewPointRecordDAO() *PointslogDAO {
	return &PointslogDAO{}
}

func (dao *PointslogDAO) CreatePointslog(ctx context.Context, log *model.Pointslog) error {
	return utils.DB.WithContext(ctx).Create(log).Error
}

// 获取指定用户积分日志
func (dao *PointslogDAO) GetPointslogByID(ctx context.Context, userID int, query map[string]any) (*[]model.Pointslog, int64, error) {
	var logs []model.Pointslog
	var total int64

	pageNum := query["page_num"].(int)
	pageSize := query["page_size"].(int)

	db := utils.DB.WithContext(ctx).Model(&model.Pointslog{}).Where("user_id = ?", userID)

	if err := db.Count(&total).Error; err != nil {
		return nil, -1, err
	}

	err := db.Offset((pageNum - 1) * pageSize).
		Limit(pageSize).
		Find(&logs).Error

	return &logs, total, err
}

// 获取积分日志
func (PointslogDAO) GetPointslogs(ctx context.Context, query map[string]any) (*[]model.Pointslog, int64, error) {
	var logs []model.Pointslog
	var total int64

	pageNum := query["page_num"].(int)
	pageSize := query["page_size"].(int)

	db := utils.DB.WithContext(ctx).Model(&model.Pointslog{})

	if err := db.Count(&total).Error; err != nil {
		return nil, -1, err
	}

	err := db.Offset((pageNum - 1) * pageSize).
		Limit(pageSize).
		Find(&logs).Error

	return &logs, total, err
}

func (dao *PointslogDAO) GetUserPointsBalance(ctx context.Context, userID int) (int, error) {
	var user model.User
	err := utils.DB.WithContext(ctx).Select("points").First(&user, userID).Error
	return 0, err
}
