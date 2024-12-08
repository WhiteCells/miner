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
	userInfoTimeout  = 30 * time.Minute
	userTokenTimeout = 24 * time.Hour
	loginIPTimeout   = 24 * time.Hour
)

func (c *UserCache) SetUserInfoByID(ctx context.Context, user *model.User) error {
	key := fmt.Sprintf("user:%d:info", user.ID)
	userJSON, err := json.Marshal(user)
	if err != nil {
		return err
	}
	return utils.RDB.Set(ctx, key, string(userJSON), userInfoTimeout)
}

func (c *UserCache) GetUserInfoByID(ctx context.Context, userID int) (*model.User, error) {
	key := fmt.Sprintf("user:%d:info", userID)
	userJSON, err := utils.RDB.Get(ctx, key)
	if err != nil {
		return nil, err
	}
	var user model.User
	err = json.Unmarshal([]byte(userJSON), &user)
	return &user, err
}

func (c *UserCache) SetUserInfoByName(ctx context.Context, user *model.User) error {
	key := fmt.Sprintf("user:%s:info", user.Name)
	userJSON, err := json.Marshal(user)
	if err != nil {
		return err
	}
	return utils.RDB.Set(ctx, key, userJSON, userInfoTimeout)
}

func (c *UserCache) GetUserInfoByName(ctx context.Context, userName string) (*model.User, error) {
	key := fmt.Sprintf("user:%s:info", userName)
	userJSON, err := utils.RDB.Get(ctx, key)
	if err != nil {
		return nil, err
	}
	var user model.User
	err = json.Unmarshal([]byte(userJSON), &user)
	return &user, err
}

func (c *UserCache) SetUserTokenByID(ctx context.Context, userID int, token string) error {
	key := fmt.Sprintf("user:%d:token", userID)
	return utils.RDB.Set(ctx, key, token, userTokenTimeout)
}

func (c *UserCache) GetUserTokenByID(ctx context.Context, userID int) (string, error) {
	key := fmt.Sprintf("user:%d:token", userID)
	return utils.RDB.Get(ctx, key)
}

func (c *UserCache) DeleteUserCache(ctx context.Context) error {
	pattern := "user:*"

	keys, err := utils.RDB.Scan(ctx, pattern)
	if err != nil {
		return fmt.Errorf("scan keys error: %w", err)
	}

	for _, key := range keys {
		if err := utils.RDB.Del(ctx, key); err != nil {
			return fmt.Errorf("failed to delete key %s: %w", key, err)
		}
	}

	return nil
}
