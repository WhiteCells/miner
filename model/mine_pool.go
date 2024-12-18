package model

type MinePool struct {
	ID   int     `json:"id" gorm:"column:id;type:int;primaryKey;autoIncrement;comment:矿池ID"`
	SSL  string  `json:"ssl" gorm:"column:ssl;type:varchar(255);comment:ssl链接"`
	Cost float64 `json:"cost" gorm:"column:cost;double:int;comment:费用"`
}

func (MinePool) TableName() string {
	return "mine_pool"
}
