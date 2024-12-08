package mysql

import (
	"miner/model"
	"miner/utils"

	"gorm.io/gorm"
)

type MinerDAO struct{}

func NewMinerDAO() *MinerDAO {
	return &MinerDAO{}
}

// 创建矿机
func (dao *MinerDAO) CreateMiner(miner *model.Miner) error {
	return utils.DB.Create(miner).Error
}

// 获取矿机信息
func (dao *MinerDAO) GetMinerByID(id int) (*model.Miner, error) {
	var miner model.Miner
	err := utils.DB.First(&miner, id).Error
	return &miner, err
}

// 获取矿场的所有矿机
func (dao *MinerDAO) GetFarmMiners(farmID int) ([]model.Miner, error) {
	var miners []model.Miner
	err := utils.DB.Joins("JOIN farm_miners ON miners.id = farm_miners.miner_id").
		Where("farm_miners.farm_id = ?", farmID).
		Find(&miners).Error
	return miners, err
}

// 更新矿机信息
func (dao *MinerDAO) UpdateMiner(miner *model.Miner) error {
	return utils.DB.Save(miner).Error
}

// 删除矿机
func (dao *MinerDAO) DeleteMiner(minerID int) error {
	return utils.DB.Transaction(func(tx *gorm.DB) error {
		// 删除矿机关联
		if err := tx.Where("miner_id = ?", minerID).Delete(&model.FarmMiner{}).Error; err != nil {
			return err
		}
		if err := tx.Where("miner_id = ?", minerID).Delete(&model.MinerFlightsheet{}).Error; err != nil {
			return err
		}
		// 删除矿机
		return tx.Delete(&model.Miner{}, minerID).Error
	})
}

// 更新矿机状态
func (dao *MinerDAO) UpdateMinerStatus(minerID int, status int) error {
	return utils.DB.
		Model(&model.Miner{}).
		Where("id = ?", minerID).
		Update("status", status).Error
}

// 转移矿机所有权
func (dao *MinerDAO) TransferMiner(minerID int, fromFarmID int, toFarmID int) error {
	return utils.DB.Transaction(func(tx *gorm.DB) error {
		// 删除原矿场关联
		if err := tx.Where("miner_id = ? AND farm_id = ?", minerID, fromFarmID).
			Delete(&model.FarmMiner{}).Error; err != nil {
			return err
		}
		// 创建新矿场关联
		return tx.Create(&model.FarmMiner{
			MinerID: minerID,
			FarmID:  toFarmID,
		}).Error
	})
}
