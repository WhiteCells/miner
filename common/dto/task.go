package dto

/*
// detail
task:<task_id>/detail

// 命令，统计，配置的结果需要存储，
user:<user_id>:<farm_id>:<miner_id>

*/

type Task struct {
	ID      string `json:"id"`
	UserID  string `json:"user_id"`
	FarmID  string `json:"farm_id"`
	MinerID string `json:"miner_id"`
	Command string `json:"command"`
	Status  string `json:"status"`
	Stats   string `json:"stats"`
}

type SendCmdReq struct {
	FarmID  string `json:"farm_id"`
	MinerID string `json:"miner_id"`
	Cmd     string `json:"cmd"`
}

type SetConfigReq struct {
	FramID  string `json:"fram_id"`
	MinerID string `json:"miner_id"`
	Config  string `json:"config"`
}

/*
Config

*/
