package service

import (
	"context"
	"fmt"
	"miner/dao/mysql"
	"miner/model"
)

type OperLogService struct {
	operLogDAO *mysql.OperLogDAO
}

func NewOperLogService() *OperLogService {
	return &OperLogService{
		operLogDAO: mysql.NewOperLogDAO(),
	}
}

// GetOperLogs 获取用户操作日志
func (s *OperLogService) GetOperLogs(ctx context.Context, query map[string]interface{}) (*[]model.Operlog, int64, error) {
	// 调用 DAO 层获取日志
	logs, total, err := s.operLogDAO.GetOperLogs(query)
	if err != nil {
		return nil, -1, fmt.Errorf("failed to get oper logs: %w", err)
	}
	return logs, total, nil
}
