package model

type Miner struct {
	ID    int    `json:"id" gorm:"column:id;type:int;primaryKey;autoIncrement;comment:矿机唯一标识"`
	Name  string `json:"name" gorm:"column:name;type:varchar(255);comment:矿机名"`
	RigID string `json:"rig_id" gorm:"index;column:rig_id;type:varchar(255);comment:rigID"`
	Pass  string `json:"pass" gorm:"column:pass;type:varchar(255);comment:rig密码"`
}

func (Miner) TableName() string {
	return "miner"
}

func GetMinerAllowChangeField() map[string]bool {
	return map[string]bool{
		"name": true,
	}
}
