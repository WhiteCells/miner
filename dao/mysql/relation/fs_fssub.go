package relation

type FsFssub struct {
	FsID  int `gorm:"index"`
	Fssub int `gorm:"index"`
}

func (FsFssub) TableName() string {
	return "fs_fssub"
}
