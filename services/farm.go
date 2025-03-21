package services

import (
	"miner/common/dto"
	"miner/dao/mysql"
	"miner/model"
	"miner/utils"

	"github.com/gin-gonic/gin"
)

type FarmService struct {
	farmDAO *mysql.FarmDAO
}

func NewFarmService() *FarmService {
	return &FarmService{
		farmDAO: mysql.NewFarmDAO(),
	}
}

func (m *FarmService) CreateFarm(ctx *gin.Context, userID int, req *dto.CreateFarmReq) error {
	farm := &model.Farm{
		Name:     req.Name,
		TimeZone: req.TimeZone,
	}
	return m.farmDAO.CreateFarm(farm, userID)
}

func (m *FarmService) DelFarm(ctx *gin.Context, userID int, req *dto.DeleteFarmReq) error {
	return m.farmDAO.DeleteFarmByID(req.FarmID, userID)
}

func (m *FarmService) UpdateFarm(ctx *gin.Context, req *dto.UpdateFarmReq) error {
	farm := &model.Farm{}
	if name, ok := req.UpdateInfo["name"].(string); ok {
		farm.Name = name
	}
	if timezone, ok := req.UpdateInfo["time_zone"].(string); ok {
		farm.TimeZone = timezone
	}
	return m.farmDAO.UpdateFarm(req.FarmID, farm)
}

func (m *FarmService) GetFarm(ctx *gin.Context, farmID int) (*model.Farm, error) {
	farm := &model.Farm{}
	err := utils.DB.Where("id=?", farmID).First(farm).Error
	return farm, err
}
