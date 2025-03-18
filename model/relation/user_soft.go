package relation

type UserSoft struct {
	UserID int `gorm:"index"`
	SoftID int `gorm:"index"`
}

func (UserSoft) TableName() string {
	return "user_soft"
}
