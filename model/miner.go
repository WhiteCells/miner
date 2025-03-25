package model

type Miner struct {
	// 使用 ID 作为 RigID，hiveos firstrun 脚本中，RigID 没有做长度限制，只需要是数字
	ID   int    `json:"id" gorm:"column:id;type:int;primaryKey;autoIncrement;comment:矿机唯一标识"`
	Name string `json:"name" gorm:"column:name;type:varchar(255);comment:矿机名"`
	Pass string `json:"pass" gorm:"column:pass;type:varchar(255);comment:rig密码"`
}

func (Miner) TableName() string {
	return "miner"
}

func GetMinerAllowChangeField() map[string]bool {
	return map[string]bool{
		"name": true,
	}
}
