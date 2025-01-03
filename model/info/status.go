package info

type MinerStatus struct {
	FsID       int       `json:"fs_id"`
	Coin       string    `json:"coin"`
	Temp       []int     `json:"temp"`
	Fan        []int     `json:"fan"`
	Power      []int     `json:"power"`
	Df         string    `json:"df"`
	Mem        []int     `json:"mem"`
	Cputemp    []int     `json:"cputemp"`
	Cpuavg     []float32 `json:"cpuavg"`
	Miner      string    `json:"miner"`
	TotalKhs   int       `json:"total_khs"`
	Status     string    `json:"status"`
	Khs        string    `json:"khs"`
	Hs         []float32 `json:"hs"`
	HsUnits    string    `json:"hs_units"`
	Algo       string    `json:"algo"`
	BusNumbers []int     `json:"bus_numbers"`
}
