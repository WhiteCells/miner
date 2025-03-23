package model

type Farm struct {
	ID       int    `json:"id" gorm:"column:id;type:int;primaryKey;autoIncrement;comment:矿场唯一标识"`
	Name     string `json:"name" gorm:"column:name;type:varchar(255);comment:矿场名"`
	TimeZone string `json:"time_zone" gorm:"column:time_zone;type:varchar(255);comment:时区"`
	Hash     string `json:"hash" gorm:"column:hash;type:varchar(255);comment:矿场Hash"`
	GpuNum   int    `json:"gpu_num" gorm:"column:gpu_num;type:int"`
	MinerNum int    `json:"miner_num" gorm:"column:miner_num;type:int"`
}

func (Farm) TableName() string {
	return "farm"
}

func GetFarmallowChangeField() map[string]bool {
	return map[string]bool{
		"name":      true,
		"coin_type": true,
		"mine_pool": true,
		"hash":      true,
	}
}
