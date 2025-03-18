package mysql

import (
	"miner/model"
	"miner/utils"
)

// 避免引用循环
func Migrate() error {
	return utils.DB.AutoMigrate(
		&model.User{},
		&model.Farm{},
		&model.Miner{},
		&model.Flightsheet{},
		&model.Wallet{},
		&model.UserFarm{},
		&model.UserWallet{},
		&model.FarmMiner{},
		&model.UserMiner{},
		&model.MinerFlightsheet{},
		&model.FlightsheetWallet{},
		&model.OperLog{},
		&model.LoginLog{},
		&model.PointsRecord{},
		&model.UserFlightsheet{},
		&model.FarmFlightsheet{},
		&model.MinePool{},
		&model.System{},
		&model.Task{},
	)
}
