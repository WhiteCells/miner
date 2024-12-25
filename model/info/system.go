package info

type System struct {
	PointsReward   int      `json:"points_reward"`
	InviteReward   int      `json:"invite_reward"`
	RegisterSwitch bool     `json:"register_switch"`
	GlobalMinePool MinePool `json:"global_mine_pool"`
}
