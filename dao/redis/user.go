package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"miner/model"
	"miner/utils"
	"time"
)

type UserCache struct{}

func NewUserCache() *UserCache {
	return &UserCache{}
}

const (
	userKeyPrefix    = "user:"
	userInfoTimeout  = 30 * time.Minute
	userTokenTimeout = 24 * time.Hour
	loginIPTimeout   = 24 * time.Hour
)

// {user:<id>:info, []}
func (c *UserCache) SetUserInfo(ctx context.Context, user *model.User) error {
	key := fmt.Sprintf("%s%d:info", userKeyPrefix, user.ID)
	userJSON, err := json.Marshal(user)
	if err != nil {
		return err
	}
	return utils.RDB.Set(ctx, key, string(userJSON), userInfoTimeout)
}

// 获取 user:<id>:info 的值
func (c *UserCache) GetUserInfo(ctx context.Context, userID int) (*model.User, error) {
	key := fmt.Sprintf("%s%d:info", userKeyPrefix, userID)
	userJSON, err := utils.RDB.Get(ctx, key)
	if err != nil {
		return nil, err
	}

	var user model.User
	err = json.Unmarshal([]byte(userJSON), &user)
	return &user, err
}

// {user:<id>:token, []}
func (c *UserCache) SetUserToken(ctx context.Context, userID int, token string) error {
	key := fmt.Sprintf("%s%d:token", userKeyPrefix, userID)
	return utils.RDB.Set(ctx, key, token, userTokenTimeout)
}

// 获取 user:<id>:token 的值
func (c *UserCache) GetUserToken(ctx context.Context, userID int) (string, error) {
	key := fmt.Sprintf("%s%d:token", userKeyPrefix, userID)
	return utils.RDB.Get(ctx, key)
}

// {user:<id>:last_ip, []}
func (c *UserCache) SetLoginIP(ctx context.Context, userID int, ip string) error {
	key := fmt.Sprintf("%s%d:last_ip", userKeyPrefix, userID)
	return utils.RDB.Set(ctx, key, ip, loginIPTimeout)
}

// 获取 user:<id>:last_ip 的值
func (c *UserCache) GetLoginIPByID(ctx context.Context, userID int) (string, error) {
	key := fmt.Sprintf("%s%d:last_ip", userKeyPrefix, userID)
	return utils.RDB.Get(ctx, key)
}

// {user:<name>:last_ip, []}
func (c *UserCache) GetLoginIPByName(ctx context.Context, userName string) (string, error) {
	key := fmt.Sprintf("%s%s:last_ip", userKeyPrefix, userName)
	return utils.RDB.Get(ctx, key)
}

// 删除缓存
func (c *UserCache) DeleteUserCache(ctx context.Context, userID int) error {
	keys := []string{
		fmt.Sprintf("%s%d:info", userKeyPrefix, userID),
		fmt.Sprintf("%s%d:token", userKeyPrefix, userID),
		fmt.Sprintf("%s%d:last_ip", userKeyPrefix, userID),
	}

	for _, key := range keys {
		if err := utils.RDB.Del(ctx, key); err != nil {
			return err
		}
	}
	return nil
}

// todo
// func (c *UserCache) createKey() {
// 	return fmt.Sprintf("%s%s%s")
// }
