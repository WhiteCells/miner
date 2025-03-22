package services

import (
	"context"
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
	return nil
}

func (m *FssubService) GetFssubByID(ctx context.Context, fssubID int) (*model.Fssub, error) {
	return m.fssubDAO.GetFssubByID(ctx, fssubID)
}

func (m *FssubService) GetFssubByFsID(ctx context.Context, fsID int, query map[string]any) (*[]model.Fssub, int64, error) {
	return m.fssubDAO.GetFssubByFsID(ctx, fsID, query)
}
