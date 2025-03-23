package model

type Fs struct {
	ID   int    `json:"id" gorm:"column:id;type:int;primaryKey;autoIncrement;comment:飞行表ID"`
	Name string `json:"name" gorm:"unqiue;column:name;type:varchar(255);comment:飞行表名"`
}

func (Fs) TableName() string {
	return "fs"
}

func GetFsAllowChangeField() map[string]bool {
	return map[string]bool{
		"name": true,
	}
}
