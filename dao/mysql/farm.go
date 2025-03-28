package mysql

import (
	"context"
	"miner/common/perm"
	"miner/model"
	"miner/model/relation"
	"miner/utils"

	"gorm.io/gorm"
)

type FarmDAO struct{}

func NewFarmDAO() *FarmDAO {
	return &FarmDAO{}
}

// 创建矿场
func (FarmDAO) CreateFarm(ctx context.Context, userID int, farm *model.Farm) error {
	return utils.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// farm
		if err := tx.Create(farm).Error; err != nil {
			return err
		}
		// user-farm
		userFarm := &relation.UserFarm{
			UserID: userID,
			FarmID: farm.ID,
			Perm:   perm.FarmOwner,
		}
		if err := tx.Create(userFarm).Error; err != nil {
			return err
		}
		return nil
	})
}

// 删除矿场
func (FarmDAO) DelFarmByID(ctx context.Context, userID, farmID int) error {
	return utils.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// user-farm
		if err := tx.Delete(&relation.UserFarm{}, "user_id = ? AND farm_id = ?", userID, farmID).Error; err != nil {
			return err
		}
		// farm-miner
		if err := tx.Delete(&relation.FarmMiner{}, "farm_id = ?", farmID).Error; err != nil {
			return err
		}
		// farm
		if err := tx.Delete(&model.Farm{}, farmID).Error; err != nil {
			return err
		}
		// todo delete miner ?
		return nil
	})
}

// 更新矿场信息
func (FarmDAO) UpdateFarm(ctx context.Context, userID, farmID int, updates map[string]any) error {
	return utils.DB.WithContext(ctx).
		Model(&model.Farm{}).
		Where("id=?", farmID).
		Updates(updates).Error
}

// 通过 hash 获取矿场
func (FarmDAO) GetFarmByHash(ctx context.Context, hash string) (*model.Farm, error) {
	var farm model.Farm
	err := utils.DB.WithContext(ctx).First(&farm, "hash=?", hash).Error
	return &farm, err
}

// 获取所有矿场
func (FarmDAO) GetFarms(ctx context.Context, query map[string]any) ([]model.Farm, int64, error) {
	var farms []model.Farm
	var total int64

	pageNum := query["page_num"].(int)
	pageSize := query["page_size"].(int)

	db := utils.DB.WithContext(ctx).Model(&model.Farm{})

	if err := db.
		Count(&total).Error; err != nil {
		return nil, -1, err
	}

	if err := db.
		Offset((pageNum - 1) * pageSize).
		Limit(pageSize).
		Find(&farms).Error; err != nil {
		return nil, -1, err
	}

	return farms, total, nil
}

// 获取用户的矿场
func (FarmDAO) GetFarmsByUserID(ctx context.Context, userID int, query map[string]any) ([]model.Farm, int64, error) {
	var farms []model.Farm
	var total int64

	pageNum := query["page_num"].(int)
	pageSize := query["page_size"].(int)

	if err := utils.DB.WithContext(ctx).
		Model(relation.UserFarm{}).
		Where("user_id = ?", userID).
		Count(&total).Error; err != nil {
		return nil, -1, err
	}

	if err := utils.DB.WithContext(ctx).
		Joins("JOIN user_farm ON farm.id = user_farm.farm_id").
		Where("user_farm.user_id = ?", userID).
		Offset((pageNum - 1) * pageSize).
		Limit(pageSize).
		Find(&farms).Error; err != nil {
		return nil, -1, err
	}

	return farms, total, nil
}

// 获取指定矿场
func (FarmDAO) GetFarmByFarmID(ctx context.Context, farmID int) (*model.Farm, error) {
	var farm model.Farm
	err := utils.DB.WithContext(ctx).First(&farm, farmID).Error
	return &farm, err
}

// 获取所有矿场
func (FarmDAO) GetAllFarmsByUserID(ctx context.Context, userID int) ([]model.Farm, error) {
	var farms []model.Farm
	err := utils.DB.WithContext(ctx).
		Model(&model.Farm{}).
		Joins("JOIN user_farm ON user_farm.farm_id=farm").
		Where("user_farm.user_id=?", userID).Error
	return farms, err
}

// 矿场应用飞行表
func (FarmDAO) ApplyFs(ctx context.Context, userID, farmID, fsID int) error {
	return utils.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 删除原有 farm-fs 关联
		if err := tx.Delete(&relation.FarmFs{}, "farm_id = ?", farmID).Error; err != nil {
			return err
		}
		// 建立新的 farm-fs 关联
		farmFlightsheet := &relation.FarmFs{
			FarmID: farmID,
			FsID:   fsID,
		}
		if err := tx.Create(farmFlightsheet).Error; err != nil {
			return err
		}
		// TODO 对矿场下没有设置飞行表的矿机的应用
		return nil
	})
}

// 转移
func (dao *FarmDAO) Transfer(ctx context.Context, userID, toUserID, farmID int) error {
	return utils.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 更新 user-farm 关联
		if err := tx.Model(&relation.UserFarm{}).
			Where("user_id = ?", userID).
			Update("user_id", toUserID).
			Error; err != nil {
			return err
		}
		// 更新 user-miner 关联
		// if err := tx.Model(&relation.UserMiner{}).
		// 	Where("user_id = ?", userID).
		// 	Updates(map[string]any{"user_id": toUserID}).
		// 	Error; err != nil {
		// 	return err
		// }
		return nil
	})
}

// 退出矿场
func (FarmDAO) QuitFarm(ctx context.Context, farmID int, userID int) error {
	// user-farm 关联
	// todo farm 有多个管理者
	return utils.DB.WithContext(ctx).
		Where("farm_id=? AND userID=?", farmID, userID).
		Delete(&relation.UserFarm{}).Error
}
