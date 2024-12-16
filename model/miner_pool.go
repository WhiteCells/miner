package model

type MinerPool struct {
	SSL  string `json:"ssl" gorm:"column:ssl;type:varchar(255);comment:ssl链接"`
	Cost int    `json:"cost" gorm:"column:cost;type:int;comment:费用"`
}
