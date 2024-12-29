package redis

var AdminSwitchRegisterField = "switch_register"
var AdminRewardInviteField = "reward_invite"
var AdminRewardRechargeField = "reward_invite"
var AdminGfsField = "gfs"

var UserField = "user"
var FarmField = "farm"
var MinerField = "miner"
var FsField = "fs"
var WalletField = "wallet"

var MpField = "mp"

var OsField = "os"

var MinerFsField = "miner_fs"
var FsWalletField = "fs_wallet"
var FsMinepoolField = "fs_minepool"

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
