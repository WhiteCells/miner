package mysql

import (
	"context"
	"miner/model"
	"miner/model/relation"
	"miner/utils"

	"gorm.io/gorm"
)

type MinerDAO struct{}

func NewMinerDAO() *MinerDAO {
	return &MinerDAO{}
}

// 创建矿机
func (dao *MinerDAO) CreateMiner(ctx context.Context, userID, farmID int, miner *model.Miner) error {
	// 创建矿机时就需要将用户与用户与矿机进行联系
	err := utils.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 创建 miner
		if err := tx.Create(miner).Error; err != nil {
			return err
		}

		// 建立 user-miner 关联
		userMiner := &relation.UserMiner{
			UserID:  userID,
			MinerID: miner.ID,
			// Perm:    perm.MinerOwner,
		}
		if err := tx.Create(userMiner).Error; err != nil {
			return err
		}

		// 建立 farm-miner 关联
		farmMiner := &relation.FarmMiner{
			FarmID:  farmID,
			MinerID: miner.ID,
		}
		if err := tx.Create(farmMiner).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// 删除矿机
func (dao *MinerDAO) DelMiner(ctx context.Context, userID, minerID int) error {
	err := utils.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 删除 user-miner 关联
		if err := tx.
			Where("user_id=? AND miner_id=?", userID, minerID).
			Delete(&relation.UserMiner{}).Error; err != nil {
			return err
		}
		// 删除 farm-miner 关联
		if err := tx.
			Where("miner_id=?", minerID).
			Delete(&relation.FarmMiner{}).Error; err != nil {
			return err
		}
		// 删除 miner-fs 关联
		if err := tx.
			Where("miner_id = ?", minerID).
			Delete(&relation.MinerFs{}).Error; err != nil {
			return err
		}
		// 删除 miner
		if err := tx.Delete(&model.Miner{}, minerID).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// 更新矿机
func (dao *MinerDAO) UpdateMiner(ctx context.Context, userID, minerID int, updateInfo map[string]any) error {
	return utils.DB.WithContext(ctx).
		Model(&model.Miner{}).
		Where("id=?", minerID).
		Updates(updateInfo).Error
}

// 更新矿机状态
func (dao *MinerDAO) UpdateMinerStatus(ctx context.Context, minerID int, status int) error {
	return utils.DB.WithContext(ctx).
		Model(&model.Miner{}).
		Where("id = ?", minerID).
		Update("status", status).Error
}

// 获取指定矿机
func (dao *MinerDAO) GetMinerByMinerID(ctx context.Context, userID, minerID int) (*model.Miner, error) {
	var total int64
	if err := utils.DB.WithContext(ctx).
		Model(&relation.UserMiner{}).
		Where("user_id=? AND miner_id=?", userID, minerID).
		Count(&total).Error; err != nil {
		return nil, err
	}
	var miner model.Miner
	err := utils.DB.WithContext(ctx).First(&miner, minerID).Error
	return &miner, err
}

// 获取指定矿场的所有矿机
func (dao *MinerDAO) GetMinersByFarmID(ctx context.Context, farmID int, query map[string]any) (*[]model.Miner, int64, error) {
	var miners []model.Miner
	var total int64

	pageNum := query["page_num"].(int)
	pageSize := query["page_size"].(int)

	if err := utils.DB.WithContext(ctx).
		Model(&relation.FarmMiner{}).
		Where("farm_id=?", farmID).
		Count(&total).Error; err != nil {
		return nil, -1, err
	}

	if err := utils.DB.WithContext(ctx).
		Joins("JOIN farm_miner ON miner.id = farm_miner.miner_id").
		Where("farm_miner.farm_id = ?", farmID).
		Offset((pageNum - 1) * pageSize).
		Limit(pageSize).
		Find(&miners).Error; err != nil {
		return nil, -1, err
	}

	return &miners, total, nil
}

// 获取所有矿机
func (MinerDAO) GetMiners(ctx context.Context, query map[string]any) (*[]model.Miner, int64, error) {
	var miners []model.Miner
	var total int64

	pageNum := query["page_num"].(int)
	pageSize := query["page_size"].(int)

	db := utils.DB.WithContext(ctx).Model(&model.Miner{})

	if err := db.Count(&total).Error; err != nil {
		return nil, -1, err
	}

	if err := db.
		Offset((pageNum - 1) * pageSize).
		Limit(pageSize).
		Find(&miners).Error; err != nil {
		return nil, -1, err
	}

	return &miners, total, nil
}

// 获取矿机
func (dao *MinerDAO) GetMiner(ctx context.Context, userID, farmID int, query map[string]any) (*[]model.Miner, int64, error) {
	var miners []model.Miner
	var total int64

	pageNum := query["page_num"].(int)
	pageSize := query["page_size"].(int)

	db := utils.DB.WithContext(ctx).
		Model(&model.Miner{}).
		Joins("JOIN farm_miner ON farm_miner.miner_id = miner.id").
		Joins("JOIN user_miner ON user_miner.miner_id = miner.id").
		Joins("JOIN user_farm ON user_farm.farm_id = farm_miner.farm_id").
		Where("user_farm.user_id = ? AND user_farm.farm_id = ? AND user_miner.user_id = ?", userID, farmID, userID)

	// 查询总数
	if err := db.Count(&total).Error; err != nil {
		return nil, -1, err
	}

	// 分页查询
	err := db.
		Limit(pageSize).
		Offset((pageNum - 1) * pageSize).
		Find(&miners).Error

	return &miners, total, err
}

// 应用飞行表
func (MinerDAO) ApplyFs(ctx context.Context, minerID, fsID int) error {
	return nil
}

// 取消应用飞行表
func (MinerDAO) UnApplyFs(ctx context.Context, minerID, fsID int) error {
	return nil
}

// 转移
func (dao *MinerDAO) Transfer(ctx context.Context, fromFarmID, minerID int, toFarmHash string) error {
	err := utils.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 通过 Farm hash 找到矿场
		var farm model.Farm
		if err := tx.
			Where("hash = ?", toFarmHash).
			First(&farm).Error; err != nil {
			return err
		}
		// 更新 farm-miner 关联
		if err := tx.
			Model(&relation.FarmMiner{}).
			Where("farm_id = ? AND miner_id = ?", fromFarmID, minerID).
			Update("farm_id", farm.ID).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

func (MinerDAO) ExistsRigID(ctx context.Context, rigID string) bool {
	var count int64
	err := utils.DB.WithContext(ctx).
		Model(&model.Miner{}).
		Count(&count).Error
	return err == nil
}
