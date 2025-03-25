package relationdao

import (
	"context"
	"miner/model/relation"
	"miner/utils"
)

type MinerFsDAO struct {
}

func NewMinerFsDAO() *MinerFsDAO {
	return &MinerFsDAO{}
}

func (MinerFsDAO) BindFsToMiner(ctx context.Context, fsID int, minerID int) error {
	userFarm := relation.MinerFs{
		MinerID: minerID,
		FsID:    fsID,
	}
	return utils.DB.WithContext(ctx).Create(userFarm).Error
}

func (MinerFsDAO) UnBindFsFromMiner(ctx context.Context, fsID int, minerID int) error {
	return utils.DB.
		Where("miner_id=? AND fs_id=?", minerID, fsID).
		Delete(&relation.MinerFs{}).Error
}

func (MinerFsDAO) GetFsIDFromMiner(ctx context.Context, minerID int) (int, error) {
	var minerFs relation.MinerFs
	err := utils.DB.WithContext(ctx).First(&minerFs, "miner_id=?", minerID).Error
	return minerFs.FsID, err
}

func (MinerFsDAO) GetMinerIDsFromFs(ctx context.Context, fsID int) ([]int, error) {
	var minerIDs []int
	err := utils.DB.WithContext(ctx).
		Model(&relation.MinerFs{}).
		Where("fs_id = ?", fsID).
		Pluck("miner_id", &minerIDs).Error
	return minerIDs, err
}
