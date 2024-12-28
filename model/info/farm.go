package info

import "miner/common/perm"

type Farm struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	TimeZone string `json:"time_zone"`
	// Hash     string    `json:"hash"` // ID 可以作为唯一标识
	Perm perm.FarmPerm `json:"perm"`
}
