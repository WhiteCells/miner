package services

import (
	"context"
	"miner/dao/mysql"
	"miner/model"
)

type LoginlogService struct {
	loginlogDAO *mysql.LoginlogDAO
}

func NewLoginlogService() *LoginlogService {
	return &LoginlogService{
		loginlogDAO: mysql.NewLoginlogDAO(),
	}
}

func (m *LoginlogService) GetLoginlogByID(ctx context.Context, userID int, query map[string]any) (*[]model.Loginlog, int64, error) {
	return m.loginlogDAO.GetLoginlogByID(ctx, userID, query)
}

func (m *LoginlogService) GetLoginlogs(ctx context.Context, query map[string]any) (*[]model.Loginlog, int64, error) {
	return m.loginlogDAO.GetLoginlogs(ctx, query)
}
