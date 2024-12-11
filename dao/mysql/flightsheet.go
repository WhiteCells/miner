package mysql

import (
	"miner/model"
	"miner/utils"

	"gorm.io/gorm"
)

type FlightsheetDAO struct{}

func NewFlightsheetDAO() *FlightsheetDAO {
	return &FlightsheetDAO{}
}

// CreateFlightSheet 创建飞行表
func (dao *FlightsheetDAO) CreateFlightSheet(fs *model.Flightsheet, userID int) error {
	err := utils.DB.Transaction(func(tx *gorm.DB) error {
		// 创建 fs
		if err := tx.Create(fs).Error; err != nil {
			return err
		}
		// 建立 user-fs 联系
		userFlightsheet := &model.UserFlightsheet{
			UserID:        userID,
			FlightsheetID: fs.ID,
		}
		if err := tx.Create(userFlightsheet).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteFlightSheet 删除飞行表
func (dao *FlightsheetDAO) DeleteFlightSheet(flightsheetID int, userID int) error {
	err := utils.DB.Transaction(func(tx *gorm.DB) error {
		// 删除 user-flightsheet 关联
		if err := tx.Where("flightsheet_id = ? ADN user_id = ?", flightsheetID, userID).Delete(&model.UserFlightsheet{}).Error; err != nil {
			return err
		}
		// 删除 miner-flightsheet 关联
		if err := tx.Where("flightsheet_id = ?", flightsheetID).Delete(&model.MinerFlightsheet{}).Error; err != nil {
			return err
		}
		// 删除 flightsheet-wallet 关联
		if err := tx.Where("flightsheet_id = ?", flightsheetID).Delete(&model.FlightsheetWallet{}).Error; err != nil {
			return err
		}
		// 删除飞行表
		if err := tx.Delete(&model.Flightsheet{}, flightsheetID).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateFlightSheet 更新飞行表
func (dao *FlightsheetDAO) UpdateFlightSheet(fs *model.Flightsheet) error {
	// TODO 如果钱包更新了，需要更新 飞行表-钱包 的关联
	return utils.DB.Save(fs).Error
}

// GetUserAllFlightsheet 获取用户的所有飞行表
func (dao *FlightsheetDAO) GetUserAllFlightsheet(userID int) (*[]model.Flightsheet, error) {
	var flightsheets []model.Flightsheet
	err := utils.DB.Joins("JOIN user_flightsheet ON user.id = user_flightsheet").
		Where("user_flightsheet.user_id = ?", userID).
		Find(&flightsheets).Error
	return &flightsheets, err
}

// GetFlightSheetByID 获取飞行表信息
func (dao *FlightsheetDAO) GetFlightSheetByID(flightsheetID int) (*model.Flightsheet, error) {
	var fs model.Flightsheet
	err := utils.DB.First(&fs, flightsheetID).Error
	return &fs, err
}

// ApplyFlightSheetToMiner 将飞行表应用到矿机
func (dao *FlightsheetDAO) ApplyFlightSheetToMiner(fsID int, minerID int) error {
	return utils.DB.Transaction(func(tx *gorm.DB) error {
		// 删除将原有的飞行表设置
		if err := tx.Model(&model.MinerFlightsheet{}).
			Where("miner_id = ? AND flightsheet_id = ?", minerID, fsID).
			Delete(&model.MinerFlightsheet{}).Error; err != nil {
			return err
		}
		// 创建新的关联或更新现有关联
		return tx.Where(model.MinerFlightsheet{
			MinerID:       minerID,
			FlightsheetID: fsID,
		}).FirstOrCreate(&model.MinerFlightsheet{}).Error
	})
}
