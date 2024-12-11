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

func (dao *WalletDAO) CreateWallet(wallet *model.Wallet, userID int) error {
	err := utils.DB.Transaction(func(tx *gorm.DB) error {
		// 创建 wallet
		if err := tx.Create(wallet).Error; err != nil {
			return err
		}
		// 建立 user-wallet 关联
		userWallet := &model.UserWallet{
			UserID:   userID,
			WalletID: wallet.ID,
		}
		if err := tx.Create(userWallet).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

func (dao *WalletDAO) DeleteWallet(walletID int, userID int) error {
	err := utils.DB.Transaction(func(tx *gorm.DB) error {
		// 删除 user-wallet 关联
		if err := tx.Where("user_id = ? AND wallet_id = ?", userID, walletID).Delete(&model.UserWallet{}).Error; err != nil {
			return err
		}

		// 删除 flightsheet-wallet 关联
		if err := tx.Where("wallet_id = ?", walletID).Delete(&model.FlightsheetWallet{}).Error; err != nil {
			return err
		}
		// 删除钱包
		if err := tx.Delete(&model.Wallet{}, walletID).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

func (dao *WalletDAO) UpdateWallet(wallet *model.Wallet) error {
	return utils.DB.Save(wallet).Error
}

func (dao *WalletDAO) GetUserAllWallet(userID int) (*[]model.Wallet, error) {
	var wallets []model.Wallet
	err := utils.DB.Joins("JOIN user_wallet ON wallet.id = user_wallet.wallet_id").
		Where("user_wallet.user_id = ?", userID).
		Find(&wallets).Error
	return &wallets, err
}

func (dao *WalletDAO) GetWalletByID(walletID int) (*model.Wallet, error) {
	var wallet model.Wallet
	err := utils.DB.First(&wallet, walletID).Error
	return &wallet, err
}

func (dao *WalletDAO) GetWalletByAddress(address string) (*model.Wallet, error) {
	var wallet model.Wallet
	err := utils.DB.Where("address = ?", address).First(&wallet).Error
	return &wallet, err
}
