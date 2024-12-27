package redis

import (
	"context"
	"encoding/json"
	"miner/common/points"
	"miner/model/info"
	"miner/utils"
)

type UserRDB struct{}

var userField = "user"

func NewUserRDB() *UserRDB {
	return &UserRDB{}
}

// 添加用户
// 更新用户
// +-------+-----------+--------+
// + field | key       | val    |
// +-------+-----------+--------+
// + user  | <user_id> | info   |
// +-------+-----------+--------+
//
// +------------------+------------+
// | key              | val        |
// +------------------+------------+
// | name_id:<name>   | <user_id>  |
// +------------------+------------+
// | email_id:<email> | <<user_id> |
// +------------------+------------+
func (c *UserRDB) Set(ctx context.Context, user *info.User) error {
	// 转为 json
	userJSON, err := json.Marshal(user)
	if err != nil {
		return err
	}

	nKey := MakeField("name_id", user.Name)
	eKey := MakeField("email_id", user.Email)

	pipe := utils.RDB.Client.TxPipeline()

	pipe.HSet(ctx, userField, user.ID, string(userJSON))
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

	nKey := MakeKey("name_id", user.Name)
	eKey := MakeKey("email_id", user.Email)

	pipe := utils.RDB.Client.TxPipeline()

	pipe.HDel(ctx, userField, user.ID)
	pipe.Del(ctx, nKey)
	pipe.Del(ctx, eKey)

	_, err = pipe.Exec(ctx)

	return err
}

// 获取用户信息
func (c *UserRDB) GetByID(ctx context.Context, userID string) (*info.User, error) {
	userJSON, err := utils.RDB.HGet(ctx, userField, userID)
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
	nKey := MakeKey("name_id", name)
	id, err := utils.RDB.Get(ctx, nKey)
	if err != nil {
		return nil, err
	}
	return c.GetByID(ctx, id)
}

// 通过邮箱获取用户信息
func (c *UserRDB) GetByEmail(ctx context.Context, email string) (*info.User, error) {
	eKey := MakeKey("email_id", email)
	id, err := utils.RDB.Get(ctx, eKey)
	if err != nil {
		return nil, err
	}
	return c.GetByID(ctx, id)
}

// 获取所用用户信息
func (c *UserRDB) GetAll(ctx context.Context) (*[]info.User, error) {
	idUser, err := utils.RDB.HGetAll(ctx, userField)
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
	eKey := MakeKey("name_id", name)
	_, err := utils.RDB.Get(ctx, eKey)
	return err == nil
}

// 是否存在邮箱
func (c *UserRDB) ExistsEmail(ctx context.Context, email string) bool {
	eKey := MakeKey("email_id", email)
	_, err := utils.RDB.Get(ctx, eKey)
	return err == nil
}

// 更新积分
func (c *UserRDB) UpdatePoints(ctx context.Context, userID string, num int, points_type points.PointsType) error {
	user, err := c.GetByID(ctx, userID)
	if err != nil {
		return err
	}

	switch points_type {
	case points.PointInvite:
		user.InvitePoints += num
	case points.PointRecharge:
		user.RechargePoints += num
	}

	return c.Set(ctx, user)
}
