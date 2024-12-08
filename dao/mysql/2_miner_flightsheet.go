package mysql

import (
	"miner/model"
	"miner/utils"
)

type MinerFlightsheetDAO struct{}

func NewMinerFlightsheetDAO() *MinerFlightsheetDAO {
	return &MinerFlightsheetDAO{}
}

func (dao *MinerFlightsheetDAO) CreateMinerFlightsheetDAO(minerFlightsheet *model.MinerFlightsheet) error {
	return utils.DB.Create(&minerFlightsheet).Error
}
