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

// 创建矿机
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

// 获取矿机信息
func (dao *MinerDAO) GetMinerByID(id int) (*model.Miner, error) {
	var miner model.Miner
	err := utils.DB.First(&miner, id).Error
	return &miner, err
}

// 获取矿场的所有矿机
func (dao *MinerDAO) GetFarmMiners(farmID int) (*[]model.Miner, error) {
	var miners []model.Miner
	err := utils.DB.Joins("JOIN farm_miner ON miner.id = farm_miner.miner_id").
		Where("farm_miner.farm_id = ?", farmID).
		Find(&miners).Error
	return &miners, err
}

// 更新矿机信息
func (dao *MinerDAO) UpdateMiner(miner *model.Miner) error {
	// TODO
	return utils.DB.Save(miner).Error
}

// 删除矿机
func (dao *MinerDAO) DeleteMiner(minerID int, farmID int, userID int) error {
	err := utils.DB.Transaction(func(tx *gorm.DB) error {
		// 删除 user-miner 关联
		if err := tx.Where("user_id = ? AND miner_id = ?", userID, minerID).Delete(&model.UserMiner{}).Error; err != nil {
			return err
		}
		// 删除 farm-miner 关联
		if err := tx.Where("miner_id = ?", minerID).Delete(&model.FarmMiner{}).Error; err != nil {
			return err
		}
		// 删除 miner-flightsheet 关联
		if err := tx.Where("miner_id = ?", minerID).Delete(&model.MinerFlightsheet{}).Error; err != nil {
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

// 更新矿机状态
func (dao *MinerDAO) UpdateMinerStatus(minerID int, status int) error {
	return utils.DB.
		Model(&model.Miner{}).
		Where("id = ?", minerID).
		Update("status", status).Error
}

// 转移矿机所有权
func (dao *MinerDAO) TransferMiner(minerID int, fromFarmID int, toFarmID int) error {
	err := utils.DB.Transaction(func(tx *gorm.DB) error {
		// 删除原矿场关联
		if err := tx.Where("miner_id = ? AND farm_id = ?", minerID, fromFarmID).
			Delete(&model.FarmMiner{}).Error; err != nil {
			return err
		}
		// 创建新矿场关联
		farmMiner := &model.FarmMiner{
			MinerID: minerID,
			FarmID:  toFarmID,
		}
		if err := tx.Create(farmMiner).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}
