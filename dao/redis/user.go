package redis

import (
	"context"
	"encoding/json"
	"math"
	"miner/common/points"
	"miner/model/info"
	"miner/utils"
	"time"
)

type UserRDB struct {
	farmRDB *FarmRDB
}

func NewUserRDB() *UserRDB {
	return &UserRDB{
		farmRDB: NewFarmRDB(),
	}
}

// 添加用户
// 更新用户
// +------------+-----------+--------+
// | field      | key       | val    |
// +------------+-----------+--------+
// | user:user  | <user_id> | info   |
// +------------+-----------+--------+
//
// +--------------------+------------+
// | key                | val        |
// +--------------------+------------+
// | user:name:<name>   | <user_id>  |
// +--------------------+------------+
// | user:email:<email> | <<user_id> |
// +--------------------+------------+
func (c *UserRDB) Set(ctx context.Context, user *info.User) error {
	userByte, err := json.Marshal(user)
	if err != nil {
		return err
	}

	nKey := MakeField(NameIDField, user.Name)
	eKey := MakeField(EmailIDField, user.Email)

	pipe := utils.RDB.Client.TxPipeline()

	pipe.HSet(ctx, UserField, user.ID, string(userByte))
	pipe.Set(ctx, nKey, user.ID, 0)
	pipe.Set(ctx, eKey, user.ID, 0)

	_, err = pipe.Exec(ctx)

	return err
}

// 删除用户
func (c *UserRDB) Del(ctx context.Context, userID string) error {
	user, err := c.GetByID(ctx, userID)
	if err != nil {
		return err
	}

	nKey := MakeField(NameIDField, user.Name)
	eKey := MakeField(EmailIDField, user.Email)

	pipe := utils.RDB.Client.TxPipeline()

	pipe.HDel(ctx, UserField, user.ID)
	pipe.Del(ctx, nKey)
	pipe.Del(ctx, eKey)

	_, err = pipe.Exec(ctx)

	return err
}

// 获取用户信息
func (c *UserRDB) GetByID(ctx context.Context, userID string) (*info.User, error) {
	userJSON, err := utils.RDB.HGet(ctx, UserField, userID)
	if err != nil {
		return nil, err
	}
	var user info.User
	err = json.Unmarshal([]byte(userJSON), &user)
	return &user, err
}

// 通过姓名获取用户信息
func (c *UserRDB) GetByName(ctx context.Context, name string) (*info.User, error) {
	// 通过 name 找到对应 ID
	nKey := MakeKey(NameIDField, name)
	id, err := utils.RDB.Get(ctx, nKey)
	if err != nil {
		return nil, err
	}
	return c.GetByID(ctx, id)
}

// 通过邮箱获取用户信息
func (c *UserRDB) GetByEmail(ctx context.Context, email string) (*info.User, error) {
	eKey := MakeKey(EmailIDField, email)
	id, err := utils.RDB.Get(ctx, eKey)
	if err != nil {
		return nil, err
	}
	return c.GetByID(ctx, id)
}

// 获取所用用户信息
func (c *UserRDB) GetAll(ctx context.Context) (*[]info.User, error) {
	idUser, err := utils.RDB.HGetAll(ctx, UserField)
	if err != nil {
		return nil, err
	}
	var users []info.User
	for id := range idUser {
		user, err := c.GetByID(ctx, id)
		if err != nil {
			return nil, err
		}
		users = append(users, *user)
	}
	return &users, err
}

// 是否存在用户名
func (c *UserRDB) ExistsName(ctx context.Context, name string) bool {
	_, err := c.GetByName(ctx, name)
	return err == nil
}

// 是否存在邮箱
func (c *UserRDB) ExistsEmail(ctx context.Context, email string) bool {
	_, err := c.GetByEmail(ctx, email)
	return err == nil
}

// 是否存在相同 ID
func (c *UserRDB) ExistsSameID(ctx context.Context, userID string) bool {
	_, err := c.GetByID(ctx, userID)
	return err == nil
}

// 更新积分
func (c *UserRDB) UpdatePoints(ctx context.Context, userID string, num float32, points_type points.PointsType) error {
	user, err := c.GetByID(ctx, userID)
	if err != nil {
		return err
	}

	switch points_type {
	case points.PointInvite:
		user.InvitePoints += num
	case points.PointRecharge:
		user.RechargePoints += num
	case points.PointSettlement:
		// 优先扣除 InvitePoints，不够的情况再扣除 RechargePoints
		if user.InvitePoints >= float32(math.Abs(float64(num))) {
			user.InvitePoints += num
		} else {
			user.RechargePoints += num + user.InvitePoints
			user.InvitePoints = 0
		}
	}

	return c.Set(ctx, user)
}

// SetLastCheckAt
func (c *UserRDB) SetLastCheckAt(ctx context.Context, userID string, t time.Time) error {
	user, err := c.GetByID(ctx, userID)
	if err != nil {
		return err
	}
	user.LastCheckAt = t
	return c.Set(ctx, user)
}

// token 黑名单，TTL 为 JWT TTL 剩余时间
// +-------------------+-----------+
// | key               | val       |
// +-------------------+-----------+
// | <token>:ban_token | ""        |
// +-------------------+-----------+
func (c *UserRDB) AddBanToken(ctx context.Context, token string) error {
	key := MakeKey(BanToken, token)
	// 解析 token
	claims, err := utils.ParseToken(token)
	if err != nil {
		return err
	}
	// 登出时间
	now := time.Now()
	// token 过期时间
	expTime := claims.ExpiresAt.Time
	if expTime.Before(now) {
		return err
	}
	// 计算 TTL
	ttl := time.Until(expTime)
	// 添加到 ban token
	return utils.RDB.Set(ctx, key, "", ttl)
}

// 判断 token 是否存在 ban token
func (c *UserRDB) ExistsBanToken(ctx context.Context, token string) bool {
	key := MakeKey(BanToken, token)
	if _, err := utils.RDB.Get(ctx, key); err != nil {
		return false
	}
	return true
}
