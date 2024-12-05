package model

import "time"

type PointsRecord struct {
	ID      int       `json:"id" gorm:"column:id;type:int;primaryKey;autoIncrement"`
	UserID  int       `json:"user_id" gorm:"index"`
	Type    string    `json:"type"`
	Amount  int       `json:"amount"`
	Balance int       `json:"balance"`
	Time    time.Time `json:"time"`
}

func (PointsRecord) TableName() string {
	return "points_record"
}
