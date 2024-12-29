package service

import (
	"context"
	"miner/dao/redis"
)

type HiveOsService struct {
	hiveOsRDB *redis.HiveOsRDB
}

func NewHiveOsService() *HiveOsService {
	return &HiveOsService{
		hiveOsRDB: redis.NewHiveOsRDB(),
	}
}

func (s *HiveOsService) Interact(ctx context.Context) error {

	return nil
}

func (s *HiveOsService) SendCmd(ctx context.Context) error {

	return nil
}
