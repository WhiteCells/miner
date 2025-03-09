package redis

// Admin
var AdminField = "admin"
var AdminSwitchRegisterField = "switch_register"
var AdminInviteRewardField = "invite_reward"
var AdminRechargeRatioField = "recharge_ratio"
var AdminGfsField = "gfs"

// User
var UserField = "user:user"
var EmailIDField = "user:email"
var FarmField = "farm"
var FarmHashField = "farm:hash"
var MinerField = "miner"
var FsField = "fs"
var WalletField = "wallet"
var MpField = "mp"

var MinerFsField = "miner:fs"
var FsWalletField = "fs:wallet"
var FsPoolField = "fs:pool"
var FsSoftField = "fs:soft"

var BanToken = "ban_token"

// Hiveos
var OsField = "os"
var OsInfoField = "os:info"
var OsStatsField = "os:stats"
var OsMinerField = "os:miner"
var TaskIDField = "task:id"
var TaskInfoField = "task:info"
var OsFarmHashField = "os:farm_hash"

// Mnemonic
var Mnemonic = "mnemonic"
var Active = "active"
var Non = "non"

// api
var ApiKeyBscField = "apikey:bsc"

var FreeGpuNumField = "free_gpu_num"

var CoinField = "coin"
var PoolField = "pool"
var PoolsField = "pools"
var SoftField = "soft"
var CustomField = "custom"

func MakeKey(str ...string) string {
	b := false
	res := ""
	for _, s := range str {
		if b {
			res += ":" + s
		} else {
			res += s
			b = true
		}
	}
	return res
}

func MakeField(str ...string) string {
	b := false
	res := ""
	for _, s := range str {
		if b {
			res += ":" + s
		} else {
			res += s
			b = true
		}
	}
	return res
}

func MakeVal(str ...string) string {
	b := false
	res := ""
	for _, s := range str {
		if b {
			res += ":" + s
		} else {
			res += s
			b = true
		}
	}
	return res
}
