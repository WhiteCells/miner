package services

import (
	"context"
	"errors"
	"miner/common/dto"
	"miner/dao/mysql"
	"miner/model"
)

type FsService struct {
	fsDAO *mysql.FsDAO
}

func NewFsService() *FsService {
	return &FsService{
		fsDAO: mysql.NewFsDAO(),
	}
}

func (m *FsService) CreateFs(ctx context.Context, userID int, req *dto.CreateFsReq) error {
	fs := &model.Fs{
		Name: req.Name,
	}
	return m.fsDAO.CreateFs(ctx, userID, fs, req.FssubIDs)
}

func (m *FsService) DelFs(ctx context.Context, userID, fsID int) error {
	return m.fsDAO.DelFs(ctx, userID, fsID)
}

func (m *FsService) UpdateFs(ctx context.Context, fsID int, updateInfo map[string]any) error {
	allow := model.GetFsAllowChangeField()
	updates := make(map[string]any)
	for key, val := range updateInfo {
		if allow[key] {
			updates[key] = val
		}
	}
	if len(updates) == 0 {
		return errors.New("no field update")
	}
	return m.fsDAO.UpdateFs(ctx, fsID, updates)
}

func (m *FsService) GetFsByFsID(ctx context.Context, fsID int) (*model.Fs, error) {
	return m.fsDAO.GetFsByFsID(ctx, fsID)
}

func (m *FsService) GetFsByUserID(ctx context.Context, userID int, query map[string]any) ([]model.Fs, int64, error) {
	return m.fsDAO.GetFsByUserID(ctx, userID, query)
}

func (m *FsService) GetFss(ctx context.Context, query map[string]any) ([]model.Fs, int64, error) {
	return m.fsDAO.GetFss(ctx, query)
}
