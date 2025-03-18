package relation

type UserFarm struct {
	UserID int    `gorm:"index"`
	FarmID int    `gorm:"index"`
	Perm   string `gorm:"column:perm;type:varchar(255)"`
}

func (UserFarm) TableName() string {
	return "user_farm"
}
