package info

import "miner/common/perm"

type Farm struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	TimeZone string `json:"time_zone"`
	// Hash     string    `json:"hash"` // ID 可以作为唯一标识
	Perm perm.Perm `json:"perm"` // 仅作返回
}

/*
	{
		"id": 101,
		"name": "name",
		"time_zone": "ASIA/SHANGHAI",
		"hash": "ejf490rj32r92iej2s",
		"miners": [1, 2]
	}
*/
