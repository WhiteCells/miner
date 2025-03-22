package services

import (
	"context"
	"miner/dao/mysql"
	"miner/model"
)

type SoftService struct {
	softDAO *mysql.SoftDAO
}

func NewSoftService() *SoftService {
	return &SoftService{
		softDAO: mysql.NewSoftDAO(),
	}
}

func (m *SoftService) CreateSoft(ctx context.Context, userID int, soft model.Soft) error {
	return nil
}
