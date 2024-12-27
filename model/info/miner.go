package info

import "miner/common/perm"

type Miner struct {
	ID   string         `json:"id"`
	Name string         `json:"name"`
	FS   string         `json:"fs"`
	Perm perm.MinerPerm `json:"perm"`
	// hiveos 中返回的信息
}
