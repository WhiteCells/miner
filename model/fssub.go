package model

type Fssub struct {
	ID int `gorm:"column:id;type:int;primaryKey;autoIncrement"`
}

func (Fssub) TableName() string {
	return "fssub"
}
