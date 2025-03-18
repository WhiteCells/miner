package utils

import (
	"fmt"
	"log"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DB     *gorm.DB
	onceDB sync.Once
)

// 初始化数据库连接
func InitDB() {
	onceDB.Do(func() {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			Config.MySQL.User,
			Config.MySQL.Password,
			Config.MySQL.Host,
			Config.MySQL.Port,
			Config.MySQL.DBName,
		)

		var err error
		DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
		if err != nil {
			Logger.Info("failed to connect database")
			log.Fatalf("failed to connect database: %s", err.Error())
		}

		sqlDB, err := DB.DB()
		if err != nil {
			Logger.Info("failed to get database instance")
			log.Fatalf("failed to get database instance: %s", err.Error())
		}

		sqlDB.SetMaxIdleConns(Config.MySQL.MaxIdleConns)
		sqlDB.SetMaxOpenConns(Config.MySQL.MaxOpenConns)

		Logger.Info("init database successfully")
	})
}
