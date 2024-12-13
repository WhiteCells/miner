package mysql

import (
	"miner/common/perm"
	"miner/model"
	"miner/utils"

	"gorm.io/gorm"
)

type MinerDAO struct{}

func NewMinerDAO() *MinerDAO {
	return &MinerDAO{}
}

// CreateMiner 创建矿机
func (dao *MinerDAO) CreateMiner(miner *model.Miner, userID int, farmID int) error {
	// 创建矿机时就需要将用户与用户与矿机进行联系
	err := utils.DB.Transaction(func(tx *gorm.DB) error {
		// 创建 miner
		if err := tx.Create(miner).Error; err != nil {
			return err
		}

		// 建立 user-miner 关联
		userMiner := &model.UserMiner{
			UserID:  userID,
			MinerID: miner.ID,
			Perm:    perm.MinerOwner,
		}
		if err := tx.Create(userMiner).Error; err != nil {
			return err
		}

		// 建立 farm-miner 关联
		farmMiner := &model.FarmMiner{
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

// GetMinerByID 通过矿机 ID 获取矿机信息
func (dao *MinerDAO) GetMinerByID(minerID int) (*model.Miner, error) {
	var miner model.Miner
	err := utils.DB.First(&miner, minerID).Error
	return &miner, err
}

// GetFarmMiners 获取矿场的所有矿机
func (dao *MinerDAO) GetFarmMiners(farmID int) (*[]model.Miner, error) {
	var miners []model.Miner
	err := utils.DB.
		Joins("JOIN farm_miner ON miner.id = farm_miner.miner_id").
		Where("farm_miner.farm_id = ?", farmID).
		Find(&miners).Error
	return &miners, err
}

// UpdateMiner 更新矿机信息
func (dao *MinerDAO) UpdateMiner(miner *model.Miner) error {
	return utils.DB.Save(miner).Error
}

// DeleteMiner 删除矿机
func (dao *MinerDAO) DeleteMiner(minerID int, farmID int, userID int) error {
	err := utils.DB.Transaction(func(tx *gorm.DB) error {
		// 删除 user-miner 关联
		if err := tx.
			Where("user_id = ? AND miner_id = ?", userID, minerID).
			Delete(&model.UserMiner{}).Error; err != nil {
			return err
		}
		// 删除 farm-miner 关联
		if err := tx.
			Where("miner_id = ?", minerID).
			Delete(&model.FarmMiner{}).Error; err != nil {
			return err
		}
		// 删除 miner-flightsheet 关联
		if err := tx.
			Where("miner_id = ?", minerID).
			Delete(&model.MinerFlightsheet{}).Error; err != nil {
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

// UpdateMinerStatus 更新矿机状态
func (dao *MinerDAO) UpdateMinerStatus(minerID int, status int) error {
	return utils.DB.
		Model(&model.Miner{}).
		Where("id = ?", minerID).
		Update("status", status).Error
}

// Transfer 转移矿机所有权
func (dao *MinerDAO) Transfer(minerID int, fromFarmID int, farmHash string) error {
	err := utils.DB.Transaction(func(tx *gorm.DB) error {
		// 通过 Farm hash 找到矿场
		var farm model.Farm
		if err := tx.
			Where("hash = ?", farmHash).
			First(&farm).Error; err != nil {
			return err
		}
		// 更新 farm-miner 关联
		if err := tx.
			Model(&model.FarmMiner{}).
			Where("farm_id = ? AND miner_id = ?", fromFarmID, minerID).
			Update("farm_id", farm.ID).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}
