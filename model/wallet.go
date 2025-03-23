package model

type Wallet struct {
	ID      int    `json:"id" gorm:"column:id;type:int;primaryKey;autoIncrement;comment:钱包ID"`
	Name    string `json:"name" gorm:"column:name;type:varchar(255);comment:钱包名"`
	Address string `json:"address" gorm:"column:address;type:varchar(255);comment:钱包地址"`
}

func (Wallet) TableName() string {
	return "wallet"
}

func GetWalletAllowChangeField() map[string]bool {
	return map[string]bool{
		"name":    true,
		"address": true,
	}
}
