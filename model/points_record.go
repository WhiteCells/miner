package model

import (
	"miner/common/points"
	"time"
)

type PointsRecord struct {
	ID      string            `json:"id" gorm:"column:id;type:int;primaryKey;autoIncrement"`
	UserID  string            `json:"user_id" gorm:"index;column:user_id;type:int;comment:用户ID"`
	Type    points.PointsType `json:"type" gorm:"column:type;type:varchar(255);comment:积分类型"`
	Amount  int               `json:"amount" gorm:"column:amount;type:int;comment:数量"`
	Balance int               `json:"balance" gorm:"column:balance;type:int;comment:余额"`
	Time    time.Time         `json:"time" gorm:"column:time;type:datetime;comment:时间"`
	Detail  string            `json:"detail" gorm:""`
}

func (PointsRecord) TableName() string {
	return "points_record"
}
