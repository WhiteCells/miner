package mysql

import (
	"miner/model"
	"miner/utils"

	"gorm.io/gorm"
)

type WalletDAO struct{}

func NewWalletDAO() *WalletDAO {
	return &WalletDAO{}
}

func (dao *WalletDAO) CreateWallet(wallet *model.Wallet) error {
	return utils.DB.Create(wallet).Error
}

func (dao *WalletDAO) GetWalletByID(id int) (*model.Wallet, error) {
	var wallet model.Wallet
	err := utils.DB.First(&wallet, id).Error
	return &wallet, err
}

func (dao *WalletDAO) GetWalletByAddress(address string) (*model.Wallet, error) {
	var wallet model.Wallet
	err := utils.DB.Where("address = ?", address).First(&wallet).Error
	return &wallet, err
}

func (dao *WalletDAO) UpdateWallet(wallet *model.Wallet) error {
	return utils.DB.Save(wallet).Error
}

func (dao *WalletDAO) DeleteWallet(id int) error {
	return utils.DB.Transaction(func(tx *gorm.DB) error {
		// 删除钱包关联
		if err := tx.Where("wallet_id = ?", id).Delete(&model.FlightsheetWallet{}).Error; err != nil {
			return err
		}
		// 删除钱包
		return tx.Delete(&model.Wallet{}, id).Error
	})
}
