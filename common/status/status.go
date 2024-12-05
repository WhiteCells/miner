package status

// 矿机状态
type MinerStatus int

const (
	MinerOn  MinerStatus = 0
	MinerOff MinerStatus = 1
)

// 用户状态
type UserStatus int

const (
	UserOn  UserStatus = 0
	UserOff UserStatus = 1
)
