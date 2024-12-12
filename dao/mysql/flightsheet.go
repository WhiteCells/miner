package mysql

import (
	"errors"
	"miner/model"
	"miner/utils"

	"gorm.io/gorm"
)

type FlightsheetDAO struct{}

func NewFlightsheetDAO() *FlightsheetDAO {
	return &FlightsheetDAO{}
}

// CreateFlightsheet 创建飞行表
func (dao *FlightsheetDAO) CreateFlightsheet(fs *model.Flightsheet, userID int) error {
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

// DeleteFlightsheet 删除飞行表
func (dao *FlightsheetDAO) DeleteFlightsheet(fsID int, userID int) error {
	err := utils.DB.Transaction(func(tx *gorm.DB) error {
		// 删除 user-flightsheet 关联
		if err := tx.Where("flightsheet_id = ? ADN user_id = ?", fsID, userID).Delete(&model.UserFlightsheet{}).Error; err != nil {
			return err
		}
		// 删除 miner-flightsheet 关联
		if err := tx.Where("flightsheet_id = ?", fsID).Delete(&model.MinerFlightsheet{}).Error; err != nil {
			return err
		}
		// 删除 flightsheet-wallet 关联
		if err := tx.Where("flightsheet_id = ?", fsID).Delete(&model.FlightsheetWallet{}).Error; err != nil {
			return err
		}
		// 删除飞行表
		if err := tx.Delete(&model.Flightsheet{}, fsID).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateFlightsheet 更新飞行表
func (dao *FlightsheetDAO) UpdateFlightsheet(fs *model.Flightsheet) error {
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

// GetFlightsheetByID 获取飞行表信息
func (dao *FlightsheetDAO) GetFlightsheetByID(fsID int) (*model.Flightsheet, error) {
	var fs model.Flightsheet
	err := utils.DB.First(&fs, fsID).Error
	return &fs, err
}

// GetFlightsheetCoinTypeByID 获取飞行表货币类型
func (dao *FlightsheetDAO) GetFlightsheetCoinTypeByID(fsID int) (string, error) {
	var fs model.Flightsheet
	err := utils.DB.First(&fs, fsID).Error
	return fs.CoinType, err
}

// ApplyFlightsheetToMiner 将飞行表应用到矿机
func (dao *FlightsheetDAO) ApplyFlightsheetToMiner(fsID int, minerID int) error {
	return utils.DB.Transaction(func(tx *gorm.DB) error {
		// 删除原有 miner-flightsheet-wallet 联系
		if err := tx.Model(&model.MinerFlightsheet{}).
			Where("miner_id = ? AND flightsheet_id = ?", minerID, fsID).
			Delete(&model.MinerFlightsheet{}).Error; err != nil {
			return err
		}
		// 建立新的 miner-flightsheet-wallet 联系
		minerFlightsheet := &model.MinerFlightsheet{
			MinerID:       minerID,
			FlightsheetID: fsID,
		}
		if err := tx.Create(minerFlightsheet).Error; err != nil {
			return err
		}
		return nil
	})
}

// ApplyWallet 飞行表应用钱包
func (dao *FlightsheetDAO) ApplyWallet(fsID int, walletID int) error {
	return utils.DB.Transaction(func(tx *gorm.DB) error {
		// 检查飞行表币种和钱包币种是否一致
		var fs model.Flightsheet
		if err := tx.First(&fs, fsID).Error; err != nil {
			return err
		}
		var wallet model.Wallet
		if err := tx.First(&wallet, walletID).Error; err != nil {
			return err
		}
		if fs.CoinType != wallet.CoinType {
			return errors.New("coin type inconsistent")
		}
		// 删除原有 flightsheet-wallet 联系
		if err := tx.Model(&model.FlightsheetWallet{}).
			Where("flightsheet_id = ? AND wallet_id = ?", fsID, walletID).
			Delete(&model.FlightsheetWallet{}).Error; err != nil {
			return err
		}
		// 建立新的 flightsheet-wallet 联系
		flightsheetWallet := &model.FlightsheetWallet{
			FlightsheetID: fsID,
			WalletID:      walletID,
		}
		if err := tx.Create(&flightsheetWallet).Error; err != nil {
			return err
		}
		return nil
	})
}
