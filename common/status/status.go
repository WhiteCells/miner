package status

// 矿机状态
type MinerStatus string

const (
	MinerOn  MinerStatus = "1"
	MinerOff MinerStatus = "0"
)

// 用户状态
type UserStatus string

const (
	UserOn  UserStatus = "1"
	UserOff UserStatus = "0"
)

// 注册状态
type RegisterStatus string

const (
	RegisterOn  RegisterStatus = "1"
	RegisterOff RegisterStatus = "0"
)
