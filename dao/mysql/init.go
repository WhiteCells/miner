package mysql

func Init() error {
	return AutoMigrate()
}
