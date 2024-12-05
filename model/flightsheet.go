package model

type Flightsheet struct {
	Id     int    `json:"id" gorm:"column:id;type:int;primaryKey;autoIncrement"`
	Name   string `json:"name" gorm:"unqiue;column:name;type:varchar(255)"`
	Config string `json:"config" gorm:"column:config;type:varchar(255)"`
}

func (Flightsheet) TableName() string {
	return "flightsheet"
}
