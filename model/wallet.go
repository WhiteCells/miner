package model

type Wallet struct {
	ID       int    `json:"id" gorm:"column:id;type:int;primaryKey;autoIncrement;comment:钱包ID"`
	Name     string `json:"name" gorm:"column:name;type:varchar(255);comment:钱包名"`
	Address  string `json:"address" gorm:"column:address;type:varchar(255);comment:钱包地址"`
	CoinType string `json:"coin_type" gorm:"column:coin_type;type:varchar(255);comment:货币类型"`
}

func (Wallet) TableName() string {
	return "wallet"
}
