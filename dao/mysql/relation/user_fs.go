package relation

type UserFs struct {
	UserID int `gorm:"index"`
	FsID   int `gorm:"index"`
}

func (UserFs) TableName() string {
	return "user_fs"
}
