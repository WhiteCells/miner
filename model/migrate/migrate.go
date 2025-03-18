package model

import (
	"miner/model"
	"miner/model/relation"
	"miner/utils"
)

func Migrate() error {
	return utils.DB.AutoMigrate(
		&model.Farm{},
		&model.Fs{},
		&model.LoginLog{},
		&model.Miner{},
		&model.Operlog{},
		&model.Pointslog{},
		&model.Pool{},
		&model.Soft{},
		&model.System{},
		&model.Task{},
		&model.User{},
		&model.Wallet{},

		&relation.FarmFs{},
		&relation.FarmMiner{},
		&relation.FsFssub{},
		&relation.MinerFs{},
		&relation.UserFarm{},
		&relation.UserFs{},
		&relation.UserMiner{},
		&relation.UserWallet{},
	)
}
