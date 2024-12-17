package model

type MinePool struct {
	SSL  string `json:"ssl" gorm:"column:ssl;type:varchar(255);comment:ssl链接"`
	Cost int    `json:"cost" gorm:"column:cost;type:int;comment:费用"`
}

func (MinePool) TableName() string {
	return "mine_pool"
}
