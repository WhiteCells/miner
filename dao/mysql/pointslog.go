package mysql

import (
	"miner/model"
	"miner/utils"
)

type PointslogDAO struct{}

func NewPointRecordDAO() *PointslogDAO {
	return &PointslogDAO{}
}

func (dao *PointslogDAO) CreatePointslog(log *model.Pointslog) error {
	return utils.DB.Create(log).Error
}

// 获取积分日志
func (dao *PointslogDAO) GetUserPointslog(query map[string]any) (*[]model.Pointslog, int64, error) {
	var logs []model.Pointslog
	var total int64

	db := utils.DB.Model(&model.Pointslog{})

	if userID, ok := query["user_id"].(int); ok {
		db = db.Where("user_id = ?", userID)
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

func (dao *PointslogDAO) GetUserPointsBalance(userID int) (int, error) {
	var user model.User
	err := utils.DB.Select("points").First(&user, userID).Error
	return 0, err
}
