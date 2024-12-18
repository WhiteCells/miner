package service

import (
	"context"
	"errors"
	"miner/common/dto"
	"miner/dao/mysql"
	"miner/model"
)

type FlightsheetService struct {
	flightsheetDAO     *mysql.FlightsheetDAO
	userFlightsheetDAO *mysql.UserFlightsheetDAO
}

func NewFlightsheetService() *FlightsheetService {
	return &FlightsheetService{
		flightsheetDAO:     mysql.NewFlightsheetDAO(),
		userFlightsheetDAO: mysql.NewUserFlightsheetDAO(),
	}
}

// CreateFlightsheet 创建飞行表
func (s *FlightsheetService) CreateFlightsheet(ctx context.Context, req *dto.CreateFlightsheetReq) (*model.Flightsheet, error) {
	userID, exists := ctx.Value("user_id").(int)
	if !exists {
		return nil, errors.New("invalid user_id in context")
	}

	flightsheet := &model.Flightsheet{
		Name:     req.Name,
		CoinType: req.CoinType,
		MinePool: req.MinePool,
		MineSoft: req.MineSoft,
	}

	if err := s.flightsheetDAO.CreateFlightsheet(flightsheet, userID); err != nil {
		return nil, errors.New("create flightsheet failed")
	}

	return flightsheet, nil
}

// DeleteFlightsheet 删除飞行表
func (s *FlightsheetService) DeleteFlightsheet(ctx context.Context, req *dto.DeleteFlightsheetReq) error {
	userID, exists := ctx.Value("user_id").(int)
	if !exists {
		return errors.New("invalid user_id in context")
	}

	if err := s.flightsheetDAO.DeleteFlightsheet(req.FlightsheetID, userID); err != nil {
		return errors.New("delete flightsheet failed")
	}

	return nil
}

// UpdateFlightsheet 更新飞行表
func (s *FlightsheetService) UpdateFlightsheet(ctx context.Context, req *dto.UpdateFlightsheetReq) error {
	_, exists := ctx.Value("user_id").(int)
	if !exists {
		return errors.New("invalid user_id in context")
	}

	// 查找飞行表
	flightsheet, err := s.flightsheetDAO.GetFlightsheetByID(req.FlightsheetID)
	if err != nil {
		return errors.New("flightsheet not found")
	}

	for key, value := range req.UpdateInfo {
		switch key {
		case "name":
			flightsheet.Name = value.(string)
		case "coin_type":
			flightsheet.CoinType = value.(string)
		case "wallet_id":
			// flightsheet.
			// 修改应用的钱包，但是如果失败需要回滚
		case "mine_pool":
			flightsheet.MinePool = value.(string)
		case "mine_soft":
			flightsheet.MineSoft = value.(string)
		}
	}

	if err := s.flightsheetDAO.UpdateFlightsheet(flightsheet); err != nil {
		return errors.New("delete flightsheet failed")
	}

	return nil
}

// GetFlightsheet 获取用户的所有飞行表
func (s *FlightsheetService) GetFlightsheet(ctx context.Context, query map[string]interface{}) (*[]model.Flightsheet, int64, error) {
	userID, exists := ctx.Value("user_id").(int)
	if !exists {
		return nil, -1, errors.New("invalid user_id in context")
	}
	user, total, err := s.flightsheetDAO.GetFlightsheet(userID, query)
	if err != nil {
		return nil, -1, errors.New("get flightsheet failed")
	}
	return user, total, err
}

// ApplyWallet 飞行表应用钱包
func (s *FlightsheetService) ApplyWallet(ctx context.Context, req *dto.ApplyFlightsheetWalletReq) error {
	if err := s.flightsheetDAO.ApplyWallet(req.FlightsheetID, req.WaleltID); err != nil {
		return err
	}
	return nil
}
