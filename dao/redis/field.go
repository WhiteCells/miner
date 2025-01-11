package redis

// Admin
var AdminField = "admin"
var AdminSwitchRegisterField = "switch_register"
var AdminInviteRewardField = "invite_reward"
var AdminRechargeRatioField = "recharge_ratio"
var AdminGfsField = "gfs"

// User
var UserField = "user"
var FarmField = "farm"
var MinerField = "miner"
var FsField = "fs"
var WalletField = "wallet"
var MpField = "mp"

var MinerFsField = "miner_fs"
var FsWalletField = "fs_wallet"
var FsMinepoolField = "fs_minepool"

var BanToken = "ban_token"

// Hiveos
var OsField = "os"
var OsInfoField = "os_info"
var OsStatsField = "os_stats"
var OsMinerField = "os_miner"
var TaskIDField = "task_id"
var TaskInfoField = "task_info"

// Mnemonic
var Mnemonic = "mnemonic"
var Active = "active"
var Non = "non"

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
