package model

import (
	"time"
)

type OperLog struct {
	ID       int       `json:"id" gorm:"column:id;type:int;primaryKey;autoIncrement;comment:日志唯一标识"`
	UserID   string    `json:"user_id" gorm:"column:user_id;type:varchar(255);comment:用户ID"`
	UserName string    `json:"user_name" gorm:"column:user_name;type:varchar(255);comment:用户名"`
	Time     time.Time `json:"time" gorm:"column:time;type:datetime;comment:操作时间"`
	Action   string    `json:"action" gorm:"column:action;type:varchar(255);comment:请求类型"`
	Target   string    `json:"target" gorm:"column:target;type:varchar(255);comment:请求目标"`
	IP       string    `json:"ip" gorm:"column:ip;type:varchar(255);comment:用户IP"`
	Agent    string    `json:"agent" gorm:"column:agent;type:varchar(255);comment:用户代理"`
	Status   int       `json:"status" gorm:"column:status;type:int;comment:请求返回状态"`
	Detail   string    `json:"detail" gorm:"column:detail;type:text;comment:请求回包"`
}

func (OperLog) TableName() string {
	return "oper_log"
}
