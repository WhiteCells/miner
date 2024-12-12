package mysql

import (
	"miner/model"
	"miner/utils"
)

type PointsRecordDAO struct{}

func NewPointRecordDAO() *PointsRecordDAO {
	return &PointsRecordDAO{}
}

func (dao *PointsRecordDAO) CreatePointsRecord(record *model.PointsRecord) error {
	return utils.DB.Create(record).Error
}

func (dao *PointsRecordDAO) GetUserPointsRecords(userID int, page, pageSize int) (*[]model.PointsRecord, int64, error) {
	var records []model.PointsRecord
	var total int64

	db := utils.DB.Model(&model.PointsRecord{}).Where("user_id = ?", userID)

	// 获取总数
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 获取分页数据
	err := db.Offset((page - 1) * pageSize).
		Limit(pageSize).
		Order("created_at DESC").
		Find(&records).Error

	return &records, total, err
}

func (dao *PointsRecordDAO) GetUserPointsBalance(userID int) (int, error) {
	var user model.User
	err := utils.DB.Select("points").First(&user, userID).Error
	return user.Points, err
}
