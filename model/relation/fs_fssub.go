package relation

type FsFssub struct {
	FsID    int `gorm:"index;column:fs_id;"`
	FssubID int `gorm:"index;column:fssub_id;"`
}

func (FsFssub) TableName() string {
	return "fs_fssub"
}
