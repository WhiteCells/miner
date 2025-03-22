package services

import (
	"context"
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

func (m *FsService) CreateFs(ctx context.Context, userID int, fs *model.Fs) error {
	return m.fsDAO.CreateFs(ctx, userID, fs)
}

func (m *FsService) DelFs(ctx context.Context, userID, fsID int) error {
	return m.fsDAO.DelFs(ctx, userID, fsID)
}

func (m *FsService) UpdateFs(ctx context.Context, fsID int, fs *model.Fs) error {
	return m.fsDAO.UpdateFs(ctx, fsID, fs)
}

func (m *FsService) GetFsByID(ctx context.Context, fsID int) (*model.Fs, error) {
	return m.fsDAO.GetFsByID(ctx, fsID)
}

func (m *FsService) GetFss(ctx context.Context, userID int, query map[string]any) (*[]model.Fs, int64, error) {
	return m.fsDAO.GetFss(ctx, userID, query)
}
