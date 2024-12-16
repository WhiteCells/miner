package model

import "miner/common/perm"

type Flightsheet struct {
	ID       int       `json:"id" gorm:"column:id;type:int;primaryKey;autoIncrement;comment:飞行表ID"`
	Name     string    `json:"name" gorm:"unqiue;column:name;type:varchar(255);comment:飞行表名"`
	CoinType string    `json:"coin_type" gorm:"column:coin_type;type:varchar(255);comment:货币类型"`
	MinePool string    `json:"mine_pool" gorm:"column:mine_pool;type:varchar(255);comment:矿池"`
	MineSoft string    `json:"mine_soft" gorm:"column:mine_soft;type:varchar(255);comment:挖矿软件"`
	Perm     perm.Perm `json:"perm" gorm:"column:perm;type:varchar(255);comment:用户/管理员"`
}

func (Flightsheet) TableName() string {
	return "flightsheet"
}
