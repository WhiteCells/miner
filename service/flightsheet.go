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

	if err := s.flightsheetDAO.CreateFlightSheet(flightsheet, userID); err != nil {
		return nil, errors.New("create flightsheet failed")
	}

	return flightsheet, nil
}

// DeleteFlightSheet 删除飞行表
func (s *FlightsheetService) DeleteFlightSheet(ctx context.Context, req *dto.DeleteFlightsheetReq) error {
	userID, exists := ctx.Value("user_id").(int)
	if !exists {
		return errors.New("invalid user_id in context")
	}

	if err := s.flightsheetDAO.DeleteFlightSheet(req.FlightsheetID, userID); err != nil {
		return errors.New("delete flightsheet failed")
	}

	return nil
}

// UpdateFlightSheet 更新飞行表
func (s *FlightsheetService) UpdateFlightSheet(ctx context.Context, req *dto.UpdateFlightsheetReq) error {
	_, exists := ctx.Value("user_id").(int)
	if !exists {
		return errors.New("invalid user_id in context")
	}

	// 查找飞行表
	flightsheet, err := s.flightsheetDAO.GetFlightSheetByID(req.FlightsheetID)
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

	if err := s.flightsheetDAO.UpdateFlightSheet(flightsheet); err != nil {
		return errors.New("delete flightsheet failed")
	}

	return nil
}

func (s *FlightsheetService) GetUserAllFlightsheet(ctx context.Context) (*[]model.Flightsheet, error) {
	userID, exists := ctx.Value("user_id").(int)
	if !exists {
		return nil, errors.New("invalid user_id in context")
	}
	user, err := s.userFlightsheetDAO.GetUserAllFlightsheet(userID)
	if err != nil {
		return nil, errors.New("get user all flightsheet failed")
	}
	return user, err
}
