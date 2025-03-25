package relationdao

import (
	"context"
	"errors"
	"miner/common/perm"
	"miner/model/relation"
	"miner/utils"
)

type UserFarmDAO struct {
}

func NewUserFarmDAO() *UserFarmDAO {
	return &UserFarmDAO{}
}

func (UserFarmDAO) BindUserToFarm(ctx context.Context, userID int, farmID int, p perm.FarmPerm) error {
	userFarm := relation.UserFarm{
		UserID: userID,
		FarmID: farmID,
		Perm:   p,
	}
	return utils.DB.Create(userFarm).Error
}

func (UserFarmDAO) UnBindUserFromFarm(ctx context.Context, userID int, farmID int) error {
	return utils.DB.
		Where("user_id = ? AND farm_id = ?", userID, farmID).
		Delete(&relation.UserFarm{}).Error
}

func (UserFarmDAO) GetPerm(ctx context.Context, userID int, farmID int) (perm.FarmPerm, error) {
	var userFarm relation.UserFarm
	if err := utils.DB.
		Where("user_id = ? AND farm_id = ?", userID, farmID).
		First(&userFarm).Error; err != nil {
		return perm.FarmNone, errors.New("no association")
	}
	return userFarm.Perm, nil
}
