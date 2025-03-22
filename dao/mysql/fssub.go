package mysql

import (
	"context"
	"miner/model"
	"miner/model/relation"
	"miner/utils"

	"gorm.io/gorm"
)

type FssubDAO struct {
}

func NewFssubDAO() *FssubDAO {
	return &FssubDAO{}
}

func (FssubDAO) CreateFssub(ctx context.Context, fsID int, fssub *model.Fssub) error {
	return utils.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// fssub
		if err := tx.Create(fssub).Error; err != nil {
			return err
		}
		// fs-fssub
		fsFssub := &relation.FsFssub{
			FsID:    fsID,
			FssubID: fssub.ID,
		}
		if err := tx.Create(fsFssub).Error; err != nil {
			return err
		}
		return nil
	})
}

func (FssubDAO) DelFssub(ctx context.Context, fsID, fssubID int) error {
	return utils.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// fs-fssub
		if err := tx.Delete(&relation.FsFssub{}, "fs_id=? AND fssub_id=?", fsID, fssubID).Error; err != nil {
			return err
		}
		// fssub
		if err := tx.Delete(&model.Fssub{}, "id=?", fssubID).Error; err != nil {
			return err
		}
		return nil
	})
}

func (FssubDAO) UpdateFssub(ctx context.Context, fssubID int, fssub *model.Fssub) error {
	return utils.DB.WithContext(ctx).
		Model(&model.Fssub{}).
		Where("fssub_id=?", fssubID).
		Updates(fssub).Error
}

func (FssubDAO) GetFssubByID(ctx context.Context, fssubID int) (*model.Fssub, error) {
	var fssub model.Fssub
	err := utils.DB.WithContext(ctx).
		First(&fssub, fssubID).Error
	return &fssub, err
}

func (FssubDAO) GetFssubByFsID(ctx context.Context, fsID int, query map[string]any) (*[]model.Fssub, int64, error) {
	var fssubs []model.Fssub
	var total int64

	pageNum := query["page_num"].(int)
	pageSize := query["page_size"].(int)

	// 获取总数
	if err := utils.DB.WithContext(ctx).
		Model(relation.FsFssub{}).
		Where("fs_id = ?", fsID).
		Count(&total).Error; err != nil {
		return nil, -1, err
	}

	err := utils.DB.WithContext(ctx).
		Joins("JOIN fs_fssub ON fs_fssub.fssub_id=fssub.id").
		Where("fs_fssub.fs_id=?", fsID).
		Offset((pageNum - 1) * pageSize).
		Limit(pageSize).
		Find(&fssubs).
		Error
	return &fssubs, total, err
}
