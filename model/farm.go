package model

type Farm struct {
	ID       int    `json:"id" gorm:"column:id;type:int;primaryKey;autoIncrement"`
	Name     string `json:"name" gorm:"column:name;type:varchar(255)"`
	TimeZone string `json:"time_zone" gorm:"column:time_zone;type:varchar(255)"`
	Hash     string `json:"hash" gorm:"column:hash;type:varchar(255)"`
}

func (Farm) TableName() string {
	return "farm"
}
