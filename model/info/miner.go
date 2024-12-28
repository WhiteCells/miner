package info

import "miner/common/perm"

type Miner struct {
	ID    string         `json:"id"`
	Name  string         `json:"name"`
	RigID string         `json:"rig_id"`
	Pass  string         `json:"pass"`
	FS    string         `json:"fs"`
	Perm  perm.MinerPerm `json:"perm"`
	// hiveos 中返回的信息
}
