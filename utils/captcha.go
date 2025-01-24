package utils

import (
	"context"
	"errors"
	"time"

	"github.com/mojocn/base64Captcha"
)

var captchaDriver = base64Captcha.NewDriverDigit(80, 240, 5, 0.7, 80)
var store = base64Captcha.DefaultMemStore

// 生成验证码
func GenerateCaptcha(ctx context.Context) (id string, b64s string, err error) {
	captcha := base64Captcha.NewCaptcha(captchaDriver, store)
	if captcha == nil {
		return "", "", errors.New("failed to initialize captcha")
	}
	id, b64s, _, err = captcha.Generate()
	if err != nil {
		return "", "", errors.New("generate captcha error")
	}
	if err := RDB.Client.Set(ctx, id, b64s, 20*time.Second).Err(); err != nil {
		return "", "", errors.New("set captcha error")
	}
	return id, b64s, nil
}

// 验证验证码
func VerifyCaptcha(ctx context.Context, captchaID string, captcahValue string) bool {
	storedValue, err := RDB.Client.Get(ctx, captchaID).Result()
	if err != nil {
		return false
	}
	if storedValue != captcahValue {
		return false
	}
	if err := RDB.Client.Del(ctx, captchaID).Err(); err != nil {
		Logger.Error("delete captcha error")
		return false
	}
	return true
}

// package utils

// import (
// 	"context"
// 	"errors"
// 	"time"

// 	"github.com/mojocn/base64Captcha"
// 	"github.com/redis/go-redis/v9"
// )

// type MemoryStore struct {
// 	client *redis.ClusterClient
// 	ctx    context.Context
// }

// func NewMemoryStore() *MemoryStore {
// 	return &MemoryStore{
// 		client: RDB.Client,
// 		ctx:    context.Background(),
// 	}
// }

// // Set 存储验证码
// func (m *MemoryStore) Set(id string, value string) error {
// 	_, err := m.client.Set(m.ctx, id, value, 20*time.Second).Result()
// 	return err
// }

// // Get 获取验证码
// func (m *MemoryStore) Get(id string, clear bool) string {
// 	value, err := m.client.Get(m.ctx, id).Result()
// 	if err == redis.Nil {
// 		return ""
// 	} else if err != nil {
// 		return ""
// 	}

// 	if clear {
// 		err := m.client.Del(m.ctx, id).Err()
// 		if err != nil {
// 			return ""
// 		}
// 	}

// 	return value
// }

// func (m *MemoryStore) Del(id string) error {
// 	return m.client.Del(m.ctx, id).Err()
// }

// func (m *MemoryStore) Verify(id string, value string, clear bool) bool {
// 	storeValue, err := m.client.Get(m.ctx, id).Result()
// 	if err != nil {
// 		return false
// 	}
// 	if storeValue != value {
// 		return false
// 	}
// 	if clear {
// 		m.client.Del(m.ctx, id)
// 	}
// 	return true
// }

// var captchaDriver = base64Captcha.NewDriverDigit(80, 240, 5, 0.7, 80)
// var store = NewMemoryStore()
// var captcha = base64Captcha.NewCaptcha(captchaDriver, store)

// // 生成验证码
// func GenerateCaptcha(ctx context.Context) (id string, b64s string, err error) {
// 	if captcha == nil {
// 		return "", "", errors.New("failed to initialize captcha")
// 	}
// 	id, b64s, _, err = captcha.Generate()
// 	if err != nil {
// 		return "", "", errors.New("generate captcha error")
// 	}
// 	return id, b64s, nil
// }

// // 验证验证码
// func VerifyCaptcha(ctx context.Context, captchaID string, captchaValue string) bool {
// 	storedValue := base64Captcha.DefaultMemStore.Get(captchaID, true)
// 	return storedValue == captchaValue
// }
