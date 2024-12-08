package service

import (
	"miner/dao/mysql"
)

type OperLogService struct {
	operLogDAO *mysql.OperLogDAO
}

func NewOperLogService() *OperLogService {
	return &OperLogService{
		operLogDAO: mysql.NewOperLogDAO(),
	}
}

func (s *OperLogService) WriteLog(userID int) {

}
