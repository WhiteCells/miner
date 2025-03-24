package services

import (
	"context"
	"errors"
	"miner/common/dto"
	"miner/dao/mysql"
	"miner/model"
	"miner/utils"
)

type FarmService struct {
	farmDAO *mysql.FarmDAO
}

func NewFarmService() *FarmService {
	return &FarmService{
		farmDAO: mysql.NewFarmDAO(),
	}
}

func (m *FarmService) CreateFarm(ctx context.Context, userID int, req *dto.CreateFarmReq) error {
	hash := utils.GenerateFarmHash(req.Name)
	farm := &model.Farm{
		Name:     req.Name,
		TimeZone: req.TimeZone,
		Hash:     hash,
	}
	return m.farmDAO.CreateFarm(ctx, farm, userID)
}

func (m *FarmService) DelFarm(ctx context.Context, userID, farmID int) error {
	return m.farmDAO.DelFarmByID(ctx, userID, farmID)
}

func (m *FarmService) UpdateFarm(ctx context.Context, userID, farmID int, updateInfo map[string]any) error {
	allow := model.GetFarmallowChangeField()
	updates := make(map[string]any)
	for key, val := range updateInfo {
		if allow[key] {
			updates[key] = val
		}
	}
	if len(updates) == 0 {
		return errors.New("没有可更新的字段")
	}
	return m.farmDAO.UpdateFarm(ctx, userID, farmID, updates)
}

func (m *FarmService) GetFarmByFarmID(ctx context.Context, farmID int) (*model.Farm, error) {
	return m.farmDAO.GetFarmByFarmID(ctx, farmID)
}

func (m *FarmService) GetFarms(ctx context.Context, query map[string]any) (*[]model.Farm, int64, error) {
	return m.farmDAO.GetFarms(ctx, query)
}

func (m *FarmService) ApplyFs(ctx context.Context, userID int, farmID int, fsID int) error {
	return m.farmDAO.ApplyFs(ctx, userID, farmID, fsID)
}

func (m *FarmService) Transfer(ctx context.Context, userID, toUserID int, farmID int) error {
	return m.farmDAO.Transfer(ctx, userID, toUserID, farmID)
}
