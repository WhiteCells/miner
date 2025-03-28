package dto

import "miner/model/info"

type PostTaskReq struct {
	FarmID  int           `json:"farm_id" binding:"required"`
	MinerID int           `json:"miner_id" binding:"required"`
	Type    info.TaskType `json:"type" binding:"required,oneof=cmd config"`
	Content string        `json:"content" binding:"required,max=40"`
}
