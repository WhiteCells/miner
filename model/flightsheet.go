package model

type Flightsheet struct {
	Id     int    `json:"id" gorm:"column:id;type:int;primaryKey;autoIncrement;comment:飞行表ID"`
	Name   string `json:"name" gorm:"unqiue;column:name;type:varchar(255);comment:飞行表名"`
	Config string `json:"config" gorm:"column:config;type:varchar(255);comment:飞行表配置"`
}

func (Flightsheet) TableName() string {
	return "flightsheet"
}
