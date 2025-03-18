package mysql

import (
	"errors"
	"miner/model"
	"miner/model/relation"
	"miner/utils"

	"gorm.io/gorm"
)

type FsDAO struct{}

func NewFsDAO() *FsDAO {
	return &FsDAO{}
}

// CreateFs 创建飞行表
func (dao *FsDAO) CreateFs(fs *model.Fs, userID int) error {
	err := utils.DB.Transaction(func(tx *gorm.DB) error {
		// 创建 fs
		if err := tx.Create(fs).Error; err != nil {
			return err
		}
		// 建立 user-fs 联系
		userFs := &relation.UserFs{
			UserID: userID,
			FsID:   fs.ID,
		}
		if err := tx.Create(userFs).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteFs 删除飞行表
func (dao *FsDAO) DeleteFs(fsID int, userID int) error {
	err := utils.DB.Transaction(func(tx *gorm.DB) error {
		// 删除 user-Fs 关联
		if err := tx.Where("Fs_id = ? ADN user_id = ?", fsID, userID).Delete(&relation.UserFs{}).Error; err != nil {
			return err
		}
		// 删除 miner-Fs 关联
		if err := tx.Where("Fs_id = ?", fsID).Delete(&relation.MinerFs{}).Error; err != nil {
			return err
		}
		// 删除 Fs-wallet 关联
		if err := tx.Where("Fs_id = ?", fsID).Delete(&relation.FsWallet{}).Error; err != nil {
			return err
		}
		// 删除飞行表
		if err := tx.Delete(&model.Fs{}, fsID).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateFs 更新飞行表
func (dao *FsDAO) UpdateFs(fs *model.Fs) error {
	return utils.DB.Save(fs).Error
}

// GetUserAllFs 获取用户的所有飞行表
func (dao *FsDAO) GetFs(userID int, query map[string]interface{}) (*[]model.Fs, int64, error) {
	var Fss []model.Fs
	var total int64

	// 获取总数
	if err := utils.DB.Model(relation.UserFs{}).Where("user_id = ?", userID).Count(&total).Error; err != nil {
		return nil, -1, err
	}

	pageNum := query["page_num"].(int)
	pageSize := query["page_size"].(int)

	err := utils.DB.
		Joins("JOIN user_Fs ON user.id = user_Fs.user_id").
		Where("user_Fs.user_id = ?", userID).
		Offset((pageNum - 1) * pageSize).
		Limit(pageSize).
		Find(&Fss).Error
	return &Fss, total, err
}

// GetFsByID 获取飞行表信息
func (dao *FsDAO) GetFsByID(fsID int) (*model.Fs, error) {
	var fs model.Fs
	err := utils.DB.First(&fs, fsID).Error
	return &fs, err
}

// GetFsCoinTypeByID 获取飞行表货币类型
func (dao *FsDAO) GetFsCoinTypeByID(fsID int) (string, error) {
	var fs model.Fs
	err := utils.DB.First(&fs, fsID).Error
	return fs.CoinType, err
}

// ApplyFsToMiner 将飞行表应用到矿机
func (dao *FsDAO) ApplyFsToMiner(fsID int, minerID int) error {
	return utils.DB.Transaction(func(tx *gorm.DB) error {
		// 删除原有 miner-Fs-wallet 联系
		if err := tx.Model(&relation.MinerFs{}).
			Where("miner_id = ? AND Fs_id = ?", minerID, fsID).
			Delete(&relation.MinerFs{}).Error; err != nil {
			return err
		}
		// 建立新的 miner-Fs-wallet 联系
		minerFs := &relation.MinerFs{
			MinerID: minerID,
			FsID:    fsID,
		}
		if err := tx.Create(minerFs).Error; err != nil {
			return err
		}
		return nil
	})
}

// ApplyWallet 飞行表应用钱包
func (dao *FsDAO) ApplyWallet(fsID int, walletID int) error {
	return utils.DB.Transaction(func(tx *gorm.DB) error {
		// 检查飞行表币种和钱包币种是否一致
		var fs model.Fs
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
		// 删除原有 Fs-wallet 联系
		if err := tx.Model(&relation.FsWallet{}).
			Where("Fs_id = ? AND wallet_id = ?", fsID, walletID).
			Delete(&relation.FsWallet{}).Error; err != nil {
			return err
		}
		// 建立新的 Fs-wallet 联系
		FsWallet := &relation.FsWallet{
			FsID:     fsID,
			WalletID: walletID,
		}
		if err := tx.Create(&FsWallet).Error; err != nil {
			return err
		}
		return nil
	})
}
