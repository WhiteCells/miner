package info

import "miner/common/perm"

type Farm struct {
	ID       string        `json:"id"`
	Name     string        `json:"name"`
	TimeZone string        `json:"time_zone"`
	Hash     string        `json:"hash"`
	Perm     perm.FarmPerm `json:"perm"`
	GpuNum   string        `json:"gpu_num"`
}
