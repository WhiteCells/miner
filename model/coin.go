package model

type Coin struct {
	ID   int    `json:"id" gorm:"column:id;type:int;primaryKey;autoIncrement"`
	Name string `json:"name" gorm:"column:name;type:varchar(255)"`
}

func (Coin) TableName() string {
	return "coin"
}
