package model

type Miner struct {
	ID   int    `json:"id" gorm:"column:id;type:int;primaryKey;autoIncrement;comment:矿机唯一标识"`
	Name string `json:"name" gorm:"column:name;type:varchar(255);comment:矿机名"`
}

func (Miner) TableName() string {
	return "miner"
}
