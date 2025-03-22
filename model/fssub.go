package model

type Fssub struct {
	ID   int    `gorm:"column:id;type:int;primaryKey;autoIncrement"`
	Name string `gorm:"column:name;type:varchar(255)"`
}

func (Fssub) TableName() string {
	return "fssub"
}
