package services

import (
	"context"
	"miner/dao/mysql"
	"miner/model"
)

type AdminService struct {
	adminDAO *mysql.AdminDAO
}

func NewAdminService() *AdminService {
	return &AdminService{
		adminDAO: mysql.NewAdminDAO(),
	}
}

func (m *AdminService) GetAllUsers(ctx context.Context, query map[string]any) (*[]model.User, int64, error) {
	return m.adminDAO.GetAllUsers(query)
}

func (m *AdminService) GetUserOperlogs(ctx context.Context, query map[string]any) (*[]model.Operlog, int64, error) {
	return m.adminDAO.GetUserOperlogs(query)
}

func (m *AdminService) GetUserPointslogs(ctx context.Context, query map[string]any) (*[]model.Pointslog, int64, error) {
	return m.adminDAO.GetUserPointslogs(query)
}

func (m *AdminService) GetUserLoginlogs(ctx context.Context, query map[string]any) (*[]model.Loginlog, int64, error) {
	return m.adminDAO.GetUserLoginlogs(query)
}

func (m *AdminService) SetSwitchRegister(ctx context.Context) error {
	return nil
}

func (m *AdminService) GetSwitchRegister(ctx context.Context) error {
	return nil
}
