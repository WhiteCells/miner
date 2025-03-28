package relationdao

import (
	"context"
	"miner/model/relation"
	"miner/utils"
)

type FarmMinerDAO struct {
}

func NewFarmMinerDAO() *FarmMinerDAO {
	return &FarmMinerDAO{}
}

func (FarmMinerDAO) ExistMiner(ctx context.Context, farmID, minerID int) error {
	var farmMiner relation.FarmMiner
	return utils.DB.WithContext(ctx).
		Find(&farmMiner, "farm_id=? AND miner_id=?", farmID, minerID).Error
}
