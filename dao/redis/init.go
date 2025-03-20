package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"miner/common/role"
	"miner/common/status"
	"miner/model/info"
	"miner/utils"
	"time"
)

func Init() error {
	return InitAdminRDB()
}

func InitAdminRDB() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	keys := map[string]any{}
	keys[MakeKey(AdminField, AdminInviteRewardField)] = 10
	keys[MakeKey(AdminField, AdminRechargeRatioField)] = 1
	keys[MakeKey(AdminField, AdminSwitchRegisterField)] = 1

	for key, value := range keys {
		set, err := utils.RDB.Client.SetNX(ctx, key, value, 0).Result()
		if err != nil {
			return err
		}
		if !set {
			fmt.Printf("Key %s already exists, skipping.\n", key)
		}
	}

	// user
	name := "admin"
	uid := "0"
	password, _ := utils.EncryptPassword("123456")
	email := "admin@admin.com"

	user := &info.User{
		ID:          uid,
		Name:        name,
		Password:    password,
		Secret:      "secret",
		Address:     "address",
		Email:       email,
		Role:        role.Admin,
		LastBalance: 0.0,
		Status:      status.UserOn,
		InviteCode:  uid,
		Key:         "key",
	}

	field := MakeField(UserField)
	key := MakeKey(uid)
	userByte, err := json.Marshal(user)
	if err != nil {
		return err
	}
	_, err = utils.RDB.Client.HSet(ctx, field, key, string(userByte)).Result()
	if err != nil {
		return err
	}

	field = MakeKey(EmailIDField, email)
	_, err = utils.RDB.Client.Set(ctx, field, uid, 0).Result()
	if err != nil {
		return err
	}

	// SetFreeGPUNum
	key = MakeKey(AdminField, FreeGpuNumField)
	if err := utils.RDB.Set(ctx, key, 2); err != nil {
		return err
	}

	return nil
}
