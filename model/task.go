package model

import "miner/model/info"

type Task struct {
	ID      string          `json:"id" gorm:"column:id;type:int;primaryKey;autoIncrement"`
	Type    info.TaskType   `json:"type" gorm:"column:type;type:varchar(16)"`
	Status  info.TaskStatus `json:"status" gorm:"column:status;type:varchar(16)"`
	Content string          `json:"content" gorm:"column:content;type:text"`
	Result  string          `json:"result" gorm:"column:result;type:text"`
}

func (Task) TableName() string {
	return "task"
}
