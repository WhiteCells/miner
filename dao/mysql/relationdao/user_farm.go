package relationdao

import (
	"miner/common/perm"
	"miner/model/relation"
	"miner/utils"
)

type UserFarmDAO struct {
}

func NewUserFarmDAO() *UserFarmDAO {
	return &UserFarmDAO{}
}

func (UserFarmDAO) Assign(userID int, farmID int, p perm.FarmPerm) error {
	userFarm := relation.UserFarm{
		UserID: userID,
		FarmID: farmID,
		Perm:   p,
	}
	return utils.DB.Create(userFarm).Error
}

func (UserFarmDAO) GetPerm(userID int, farmID int) (perm.FarmPerm, error) {
	var userFarm relation.UserFarm
	err := utils.DB.Where("user_id = ? AND farmID = ?", userID, farmID).First(&userFarm).Error
	return userFarm.Perm, err
}
