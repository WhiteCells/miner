package model

import (
	"miner/common/points"
	"time"
)

type Pointslog struct {
	ID      int               `json:"id" gorm:"column:id;type:int;primaryKey;autoIncrement;comment:积分日志唯一标识"`
	UserID  int               `json:"user_id" gorm:"index;column:user_id;type:int;comment:用户ID"`
	Type    points.PointsType `json:"type" gorm:"column:type;type:varchar(16);comment:积分类型"`
	Amount  float32           `json:"amount" gorm:"column:amount;type:float;comment:数量"`
	Balance float32           `json:"balance" gorm:"column:balance;type:float;comment:余额"`
	Time    time.Time         `json:"time" gorm:"column:time;type:datetime;comment:时间"`
	Detail  string            `json:"detail" gorm:"column:detail;type:varchar(255);comment:详情"`
}

func (Pointslog) TableName() string {
	return "pointslog"
}
