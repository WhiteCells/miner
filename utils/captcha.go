package utils

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/mojocn/base64Captcha"
)

var captchaDriver = base64Captcha.NewDriverDigit(80, 240, 5, 0.7, 80)
var store = base64Captcha.DefaultMemStore
var captcha = base64Captcha.NewCaptcha(captchaDriver, store)

// GenerateCaptcha
func GenerateCaptcha(ctx context.Context) (id string, b64s string, err error) {
	if captcha == nil {
		return "", "", errors.New("failed to initialize captcha")
	}
	id, b64s, answer, err := captcha.Generate()
	if err != nil {
		return "", "", errors.New("generate captcha error")
	}
	key := fmt.Sprintf("captcha:%s:%s", id, answer)
	if err := RDB.Client.Set(ctx, key, "", 200*time.Second).Err(); err != nil {
		return "", "", errors.New("set captcha error")
	}
	return id, b64s, nil
}

// VerifyCaptcha
func VerifyCaptcha(ctx context.Context, captchaID string, captchaValue string) bool {
	key := fmt.Sprintf("captcha:%s:%s", captchaID, captchaValue)
	exists, err := RDB.Client.Exists(ctx, key).Result()
	if err != nil || exists == 0 {
		return false
	}
	if err := RDB.Client.Del(ctx, key).Err(); err != nil {
		Logger.Error("delete captcha error")
		return false
	}
	return true
}
