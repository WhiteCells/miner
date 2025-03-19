package mysql

import (
	"miner/model"
	"miner/utils"
)

type PointsRecordDAO struct{}

func NewPointRecordDAO() *PointsRecordDAO {
	return &PointsRecordDAO{}
}

// CreatePointsRecord 创建积分记录
func (dao *PointsRecordDAO) CreatePointsRecord(record *model.Pointslog) error {
	return utils.DB.Create(record).Error
}

// GetUserPointsRecords 获取用户积分记录
func (dao *PointsRecordDAO) GetUserPointsRecords(query map[string]interface{}) (*[]model.Pointslog, int64, error) {
	var records []model.Pointslog
	var total int64

	db := utils.DB.Model(&model.Pointslog{})

	// 添加查询条件
	if userID, ok := query["user_id"].(int); ok {
		db = db.Where("user_id = ?", userID)
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
		Find(&records).Error

	return &records, total, err
}

func (dao *PointsRecordDAO) GetUserPointsBalance(userID int) (int, error) {
	var user model.User
	err := utils.DB.Select("points").First(&user, userID).Error
	return 0, err
}
