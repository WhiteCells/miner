package relation

type UserFarm struct {
	UserID int `gorm:"index"`
	FarmID int `gorm:"index"`
}

func (UserFarm) TableName() string {
	return "user_farm"
}
