package mysql

import (
	"context"
	"miner/model"
	"miner/model/relation"
	"miner/utils"

	"gorm.io/gorm"
)

type WalletDAO struct{}

func NewWalletDAO() *WalletDAO {
	return &WalletDAO{}
}

// 添加钱包
func (WalletDAO) CreateWallet(ctx context.Context, userID int, wallet *model.Wallet) error {
	err := utils.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 创建 wallet
		if err := tx.Create(wallet).Error; err != nil {
			return err
		}
		// 建立 user-wallet 关联
		userWallet := &relation.UserWallet{
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

// 删除钱包
func (WalletDAO) DelWallet(ctx context.Context, userID int, walletID int) error {
	err := utils.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 删除 user-wallet 关联
		if err := tx.Where("user_id = ? AND wallet_id = ?", userID, walletID).Delete(&relation.UserWallet{}).Error; err != nil {
			return err
		}

		// 删除 flightsheet-wallet 关联
		if err := tx.Where("wallet_id = ?", walletID).Delete(&relation.FsWallet{}).Error; err != nil {
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

// 更新钱包
func (WalletDAO) UpdateWallet(ctx context.Context, wallet *model.Wallet) error {
	return utils.DB.WithContext(ctx).Save(wallet).Error
}

// 获取指定钱包
func (WalletDAO) GetWalletByID(ctx context.Context, walletID int) (*model.Wallet, error) {
	var wallet model.Wallet
	err := utils.DB.WithContext(ctx).First(&wallet, walletID).Error
	return &wallet, err
}

// 获取用户所有钱包
func (WalletDAO) GetWallets(ctx context.Context, userID int, query map[string]any) (*[]model.Wallet, int64, error) {
	var wallets []model.Wallet
	var total int64

	pageNum := query["page_num"].(int)
	pageSize := query["page_size"].(int)

	// 查询总数
	if err := utils.DB.WithContext(ctx).Model(relation.UserWallet{}).Where("user_id = ?", userID).Count(&total).Error; err != nil {
		return nil, -1, err
	}
	err := utils.DB.WithContext(ctx).
		Joins("JOIN user_wallet ON user_wallet.wallet_id=wallet.id").
		Where("user.id=?", userID).
		Offset((pageNum - 1) * pageSize).
		Limit(pageSize).
		Find(&wallets).
		Error
	return &wallets, total, err
}
