package mysql

import (
	"miner/model"
	"miner/utils"
)

type UserWalletDAO struct{}

func NewUserWalletDAO() *UserWalletDAO {
	return &UserWalletDAO{}
}

func (dao *UserWalletDAO) CreateUserWalletDAO(userWallet *model.UserWallet) error {
	return utils.DB.Create(userWallet).Error
}

func (dao *UserWalletDAO) DeleteUserWalletDAO(userID int, walletID int) error {
	return utils.DB.
		Where("user_id = ? AND wallet_id = ?", userID, walletID).
		Delete(model.UserWallet{}).
		Error
}
