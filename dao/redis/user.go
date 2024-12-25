package redis

import (
	"context"
	"encoding/json"
	"miner/model/info"
	"miner/utils"
)

type UserRDB struct{}

var userField = "user"

func NewUserCache() *UserRDB {
	return &UserRDB{}
}

// 添加 User
// 更新 User
// 更新 User id->name，将用户 name 映射到 用户 ID
// +---------+-----------+-------+
// | field   |    key    |  val  |
// +---------+-----------+-------+
// | user    | <user_id> |  info |
// +---------+-----------+-------+
// | name_id |   <name>  | <id>  |
// +---------+-----------+-------+
func (c *UserRDB) Set(ctx context.Context, user *info.User) error {
	key := GenHField(UserField, user.ID)

	// 转为 json
	userJSON, err := json.Marshal(user)
	if err != nil {
		return err
	}

	//
	if err = utils.RDB.HSet(ctx, userField, key, string(userJSON)); err != nil {
		return err
	}

	// 将用户 name 映射到用户 ID
	if err = utils.RDB.Set(ctx, user.Name, user.ID); err != nil {
		return err
	}

	// 将用户邮箱 email 映射到用户 ID
	if err = utils.RDB.Set(ctx, user.Email, user.ID); err != nil {
		return err
	}

	return nil
}

// 删除用户信息
func (c *UserRDB) Del(ctx context.Context, userID string) error {
	key := GenHField(userField, userID)

	user, err := c.GetByID(ctx, userID)
	if err != nil {
		return err
	}

	pipe := utils.RDB.Client.TxPipeline()

	// 删除用户信息
	pipe.HDel(ctx, userField, key)
	// 删除用户 name-id 关联
	pipe.Del(ctx, user.Name)

	_, err = pipe.Exec(ctx)

	return err
}

// 获取用户信息
func (c *UserRDB) GetByID(ctx context.Context, userID string) (*info.User, error) {
	key := GenHField(UserField, userID)
	userJSON, err := utils.RDB.HGet(ctx, userField, key)
	if err != nil {
		return nil, err
	}
	var user info.User
	err = json.Unmarshal([]byte(userJSON), &user)
	return &user, err
}

// 通过姓名获取用户信息
func (c *UserRDB) GetByName(ctx context.Context, userName string) (*info.User, error) {
	// 通过 name-id 找到对应 ID
	id, err := utils.RDB.Get(ctx, userName)
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
	for _, userJSON := range idUser {
		var user info.User
		err = json.Unmarshal([]byte(userJSON), &user)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return &users, err
}
