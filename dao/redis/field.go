package redis

import "fmt"

var UserField = "user"
var FarmField = "farm"
var MinerField = "miner"
var FsField = "fs"
var WalletField = "wallet"

var UserFarmField = "user_farm"
var FarmMinerField = "farm_miner"

func GenHField(str1 string, str2 string) string {
	return fmt.Sprintf("%s:%s", str1, str2)
}
