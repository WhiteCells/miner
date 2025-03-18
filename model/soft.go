package model

type Soft struct {
	ID int `gorm:"column:id;type:int;primaryKey;autoIncrement"`
}

func (Soft) TableName() string {
	return "soft"
}
