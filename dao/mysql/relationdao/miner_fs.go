package relationdao

import (
	"miner/model/relation"
	"miner/utils"
)

type MinerFsDAO struct {
}

func NewMinerFsDAO() *MinerFsDAO {
	return &MinerFsDAO{}
}

func (MinerFsDAO) BindMinerToFs(minerID int, fsID int) error {
	userFarm := relation.MinerFs{
		MinerID: minerID,
		FsID:    fsID,
	}
	return utils.DB.Create(userFarm).Error
}

func (MinerFsDAO) UnBindMinerFromFs(minerID int, fsID int) error {
	return utils.DB.
		Where("miner_id=? AND fs_id=?", minerID, fsID).
		Delete(&relation.MinerFs{}).Error
}
