package points

// 获取积分的方式
type PointsType string

const (
	PointTransfer   PointsType = "transfer"
	PointRecharge   PointsType = "recharge"
	PointInvite     PointsType = "invite"
	PointSettlement PointsType = "settlement"
)
