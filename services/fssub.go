package services

import (
	"context"
	"errors"
	"miner/dao/mysql"
	"miner/model"
)

type FssubService struct {
	fssubDAO *mysql.FssubDAO
}

func NewFssubService() *FssubService {
	return &FssubService{
		fssubDAO: mysql.NewFssubDAO(),
	}
}

func (m *FssubService) CreateFssub(ctx context.Context, fsID int, fssub *model.Fssub) error {
	return m.fssubDAO.CreateFssub(ctx, fsID, fssub)
}

func (m *FssubService) DelFssub(ctx context.Context, fsID, fssubID int) error {
	return m.fssubDAO.DelFssub(ctx, fsID, fssubID)
}

func (m *FssubService) UpdateFssub(ctx context.Context, fsID, fssubID int, updateInfo map[string]any) error {
	allow := model.GetFssubAllowChangeField()
	updates := make(map[string]any)
	for key, val := range updateInfo {
		if allow[key] {
			updates[key] = val
		}
	}
	if len(updates) == 0 {
		return errors.New("no field update")
	}
	return m.fssubDAO.UpdateFssub(ctx, fssubID, updates)
}

func (m *FssubService) GetFssubByID(ctx context.Context, fssubID int) (*model.Fssub, error) {
	return m.fssubDAO.GetFssubByID(ctx, fssubID)
}

func (m *FssubService) GetFssubByFsID(ctx context.Context, fsID int, query map[string]any) (*[]model.Fssub, int64, error) {
	return m.fssubDAO.GetFssubByFsID(ctx, fsID, query)
}
