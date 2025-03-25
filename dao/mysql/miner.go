package mysql

import (
	"context"
	"errors"
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
func (dao *MinerDAO) CreateMiner(ctx context.Context, farmID int, miner *model.Miner) error {
	return utils.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// miner
		if err := tx.Create(miner).Error; err != nil {
			return errors.New("failed to create miner")
		}
		// farm-miner
		farmMiner := &relation.FarmMiner{
			FarmID:  farmID,
			MinerID: miner.ID,
		}
		if err := tx.Create(farmMiner).Error; err != nil {
			return errors.New("failed to create farm-miner")
		}
		return nil
	})
}

// 删除矿机
func (dao *MinerDAO) DelMiner(ctx context.Context, userID, minerID int) error {
	return utils.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// farm-miner
		if err := tx.
			Delete(&relation.FarmMiner{}, "miner_id=?", minerID).Error; err != nil {
			return errors.New("failed to delete farm-miner")
		}
		// miner-fs
		if err := tx.
			Delete(&relation.MinerFs{}, "miner_id=?", minerID).Error; err != nil {
			return errors.New("failed to delete miner-fs")
		}
		// miner
		if err := tx.Delete(&model.Miner{}, minerID).Error; err != nil {
			return errors.New("failed to delete miner")
		}
		return nil
	})
}

// 更新矿机
func (dao *MinerDAO) UpdateMiner(ctx context.Context, userID, minerID int, updateInfo map[string]any) error {
	return utils.DB.WithContext(ctx).
		Model(&model.Miner{}).
		Where("id=?", minerID).
		Updates(updateInfo).Error
}

// 获取指定矿机
func (dao *MinerDAO) GetMinerByID(ctx context.Context, minerID int) (*model.Miner, error) {
	var miner model.Miner
	if err := utils.DB.WithContext(ctx).First(&miner, minerID).Error; err != nil {
		return nil, errors.New("db miner not found")
	}
	return &miner, nil
}

// 获取指定矿场的所有矿机
func (dao *MinerDAO) GetMinersByFarmID(ctx context.Context, farmID int, query map[string]any) ([]model.Miner, int64, error) {
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

	return miners, total, nil
}

// 获取所有矿机
func (MinerDAO) GetMiners(ctx context.Context, query map[string]any) ([]model.Miner, int64, error) {
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

	return miners, total, nil
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
