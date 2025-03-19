package model

import "time"

type Loginlog struct {
	ID     int       `json:"id" gorm:"column:id;type:int;primaryKey;autoIncrement;comment:登陆日志唯一标识"`
	UserID string    `json:"user_id" gorm:"index;column:user_id;type:varchar(255);comment:登陆用户ID"`
	Time   time.Time `json:"time" gorm:"column:time;type:datetime;comment:登陆时间"`
	IP     string    `json:"ip" gorm:"column:ip;type:varchar(64);comment:登陆IP"`
	Status int       `json:"status" gorm:"column:status;type:int;comment:登陆状态码"`
}

func (Loginlog) TableName() string {
	return "loginlog"
}
