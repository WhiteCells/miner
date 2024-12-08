package model

type Farm struct {
	ID       int    `json:"id" gorm:"column:id;type:int;primaryKey;autoIncrement;comment:矿场唯一标识"`
	Name     string `json:"name" gorm:"column:name;type:varchar(255);comment:矿场名"`
	TimeZone string `json:"time_zone" gorm:"column:time_zone;type:varchar(255);comment:时区"`
	Hash     string `json:"hash" gorm:"column:hash;type:varchar(255);comment:矿场Hash"`
}

func (Farm) TableName() string {
	return "farm"
}
