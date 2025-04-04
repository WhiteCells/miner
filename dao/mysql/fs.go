package mysql

import (
	"context"
	"miner/model"
	"miner/model/relation"
	"miner/utils"

	"gorm.io/gorm"
)

type FsDAO struct{}

func NewFsDAO() *FsDAO {
	return &FsDAO{}
}

// 创建飞行表
func (FsDAO) CreateFs(ctx context.Context, userID int, fs *model.Fs, fssubIDs []int) error {
	err := utils.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// fs
		if err := tx.Create(fs).Error; err != nil {
			return err
		}
		// user-fs
		userFs := &relation.UserFs{
			UserID: userID,
			FsID:   fs.ID,
		}
		if err := tx.Create(userFs).Error; err != nil {
			return err
		}
		// fs-fssub
		var fsFssubs []relation.FsFssub
		for _, fssubID := range fssubIDs {
			fsFssubs = append(fsFssubs, relation.FsFssub{
				FsID:    fs.ID,
				FssubID: fssubID,
			})
		}
		if err := tx.Create(&fsFssubs).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// 删除飞行表
func (FsDAO) DelFs(ctx context.Context, userID int, fsID int) error {
	err := utils.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
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

// 更新飞行表
func (FsDAO) UpdateFs(ctx context.Context, fsID int, updates map[string]any) error {
	return utils.DB.WithContext(ctx).
		Model(&model.Fs{}).
		Where("id=?", fsID).
		Updates(updates).Error
}

// 获取指定飞行表
func (FsDAO) GetFsByFsID(ctx context.Context, fsID int) (*model.Fs, error) {
	var fs model.Fs
	err := utils.DB.WithContext(ctx).First(&fs, fsID).Error
	return &fs, err
}

// 获取指定用户指定飞行表
func (FsDAO) GetFsByUserID(ctx context.Context, userID int, query map[string]any) ([]model.Fs, int64, error) {
	var fss []model.Fs
	var total int64

	pageNum := query["page_num"].(int)
	pageSize := query["page_size"].(int)

	if err := utils.DB.WithContext(ctx).
		Model(relation.UserFs{}).
		Where("user_id = ?", userID).
		Count(&total).Error; err != nil {
		return nil, -1, err
	}

	if err := utils.DB.WithContext(ctx).
		Joins("JOIN user_fs ON user_fs.user_id=user.id").
		Where("user_fs.user_id=?", userID).
		Offset((pageNum - 1) * pageSize).
		Limit(pageSize).
		Find(&fss).
		Error; err != nil {
		return nil, -1, err
	}

	return fss, total, nil
}

// 获取所有飞行表
func (FsDAO) GetFss(ctx context.Context, query map[string]any) ([]model.Fs, int64, error) {
	var fss []model.Fs
	var total int64

	pageNum := query["page_num"].(int)
	pageSize := query["page_size"].(int)

	db := utils.DB.WithContext(ctx).Model(&model.Fs{})

	if err := db.
		Count(&total).Error; err != nil {
		return nil, -1, err
	}

	if err := db.
		Offset((pageNum - 1) * pageSize).
		Limit(pageSize).
		Find(&fss).
		Error; err != nil {
		return nil, -1, err
	}

	return fss, total, nil
}

// 矿机应用飞行表
func (FsDAO) ApplyFsToMiner(ctx context.Context, fsID int, minerID int) error {
	return utils.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
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
