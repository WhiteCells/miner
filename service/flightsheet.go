package service

import (
	"context"
	"errors"
	"miner/common/dto"
	"miner/dao/redis"
	"miner/model/info"
	"miner/utils"
)

type FlightsheetService struct {
	fsRDB *redis.FsRDB
}

func NewFlightsheetService() *FlightsheetService {
	return &FlightsheetService{
		fsRDB: redis.NewFsRDB(),
	}
}

// CreateFlightsheet 创建飞行表
func (s *FlightsheetService) CreateFlightsheet(ctx context.Context, req *dto.CreateFsReq) (*info.Fs, error) {
	userID, exists := ctx.Value("user_id").(string)
	if !exists {
		return nil, errors.New("invalid user_id in context")
	}

	id, err := utils.GenerateUID()
	if err != nil {
		return nil, err
	}
	flightsheet := &info.Fs{
		ID:   id,
		Name: req.Name,
		Coin: req.Coin,
		Pool: req.Pool,
		Soft: req.Soft,
	}

	if err := s.fsRDB.Set(ctx, userID, flightsheet); err != nil {
		return nil, errors.New("create flightsheet failed")
	}

	return flightsheet, nil
}

// DeleteFlightsheet 删除飞行表
func (s *FlightsheetService) DeleteFlightsheet(ctx context.Context, req *dto.DeleteFsReq) error {
	userID, exists := ctx.Value("user_id").(string)
	if !exists {
		return errors.New("invalid user_id in context")
	}

	if err := s.fsRDB.Del(ctx, userID, req.FsID); err != nil {
		return errors.New("delete flightsheet failed")
	}

	return nil
}

// UpdateFlightsheet 更新飞行表
func (s *FlightsheetService) UpdateFlightsheet(ctx context.Context, req *dto.UpdateFsReq) error {
	userID, exists := ctx.Value("user_id").(string)
	if !exists {
		return errors.New("invalid user_id in context")
	}

	// 查找飞行表
	flightsheet, err := s.fsRDB.GetByID(ctx, userID, req.FsID)
	if err != nil {
		return errors.New("flightsheet not found")
	}

	for key, value := range req.UpdateInfo {
		switch key {
		case "name":
			flightsheet.Name = value.(string)
		case "coin_id":
			flightsheet.Coin = value.(string)
		case "mine_id":
			flightsheet.Pool = value.(string)
		case "soft_id":
			flightsheet.Soft = value.(string)
		}
	}

	if err := s.fsRDB.Set(ctx, userID, flightsheet); err != nil {
		return errors.New("update flightsheet failed")
	}

	return nil
}

// GetFlightsheet 获取用户的所有飞行表
func (s *FlightsheetService) GetFlightsheet(ctx context.Context) (*[]info.Fs, error) {
	userID, exists := ctx.Value("user_id").(string)
	if !exists {
		return nil, errors.New("invalid user_id in context")
	}
	user, err := s.fsRDB.GetAll(ctx, userID)
	if err != nil {
		return nil, errors.New("get flightsheet failed")
	}
	return user, err
}

// ApplyWallet 飞行表应用钱包
func (s *FlightsheetService) ApplyWallet(ctx context.Context, req *dto.ApplyWalletReq) error {
	userID, exists := ctx.Value("user_id").(string)
	if !exists {
		return errors.New("invalid user_id in context")
	}
	if err := s.fsRDB.ApplyWallet(ctx, userID, req.FsID, req.WaleltID); err != nil {
		return err
	}
	return nil
}
