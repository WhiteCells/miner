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

// 创建飞行表
func (dao *FlightsheetDAO) CreateFlightSheet(fs *model.Flightsheet) error {
	return utils.DB.Create(fs).Error
}

// 获取飞行表信息
func (dao *FlightsheetDAO) GetFlightSheetByID(id int) (*model.Flightsheet, error) {
	var fs model.Flightsheet
	err := utils.DB.First(&fs, id).Error
	return &fs, err
}

// 更新飞行表
func (dao *FlightsheetDAO) UpdateFlightSheet(fs *model.Flightsheet) error {
	return utils.DB.Save(fs).Error
}

// 删除飞行表
func (dao *FlightsheetDAO) DeleteFlightSheet(id int) error {
	return utils.DB.Transaction(func(tx *gorm.DB) error {
		// 删除关联
		if err := tx.Where("flight_sheet_id = ?", id).Delete(&model.MinerFlightsheet{}).Error; err != nil {
			return err
		}
		if err := tx.Where("flight_sheet_id = ?", id).Delete(&model.FlightsheetWallet{}).Error; err != nil {
			return err
		}
		// 删除飞行表
		return tx.Delete(&model.Flightsheet{}, id).Error
	})
}

// 将飞行表应用到矿机
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
