package status

// 矿机状态
type MinerStatus int

const (
	MinerOn  MinerStatus = 1
	MinerOff MinerStatus = 0
)

// 用户状态
type UserStatus int

const (
	UserOn  UserStatus = 1
	UserOff UserStatus = 0
)

// 注册状态
type RegisterStatus int

const (
	RegisterOn  RegisterStatus = 1
	RegisterOff RegisterStatus = 0
)
