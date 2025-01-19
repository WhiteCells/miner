package dto

import "miner/model/info"

type PostTaskReq struct {
	FarmID  string        `json:"farm_id" binding:"required,max=20"`
	MinerID string        `json:"miner_id" binding:"required,max=20"`
	RigID   string        `json:"rig_id" binding:"required,max=8"`
	Type    info.TaskType `json:"type" binding:"required,oneof=cmd config"`
	Content string        `json:"content" binding:"required,max=20"`
}
