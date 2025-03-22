package services

import (
	"context"
	"miner/dao/mysql"
	"miner/model"
)

type MinerService struct {
	minerDAO *mysql.MinerDAO
}

func NewMinerService() *MinerService {
	return &MinerService{
		minerDAO: mysql.NewMinerDAO(),
	}
}

func (m *MinerService) CreateMiner(ctx context.Context, userID, farmID int, miner *model.Miner) error {
	return m.minerDAO.CreateMiner(ctx, userID, farmID, miner)
}

func (m *MinerService) DelMiner(ctx context.Context, userID, minerID int) error {
	return m.minerDAO.DelMiner(ctx, userID, minerID)
}

func (m *MinerService) UpdateMiner(ctx context.Context, miner *model.Miner) error {
	return m.minerDAO.UpdateMiner(ctx, miner)
}

func (m *MinerService) GetMinerByID(ctx context.Context, userID, minerID int) (*model.Miner, error) {
	return m.minerDAO.GetMinerByID(ctx, userID, minerID)
}

func (m *MinerService) GetMiners(ctx context.Context, farmID int, query map[string]any) (*[]model.Miner, int64, error) {
	return m.minerDAO.GetMiners(ctx, farmID, query)
}

func (m *MinerService) ApplyFs(ctx context.Context, userID, farmID, minerID, fsID int) error {
	return m.minerDAO.ApplyFs(ctx, minerID, fsID)
}
