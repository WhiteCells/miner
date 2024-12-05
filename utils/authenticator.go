package utils

import (
	"encoding/base64"
	"fmt"
	"github.com/suanju/googleAuthenticator"
	"os"
	"strings"
)

// CreateSecret 创建密钥
func CreateSecret() string {
	authenticator := googleAuthenticator.NewGoogleAuthenticator(6)

	// 创建一个 64 字节的随机密钥
	secret, err := authenticator.CreateSecret(64)
	if err != nil {
		fmt.Println("创建密钥时出错:", err)
		//return
	}
	return secret
}

// GetCode 根据密钥获取验证码
func GetCode(secret string) string {
	authenticator := googleAuthenticator.NewGoogleAuthenticator(6)
	code, err := authenticator.GetCode(secret, 0)
	if err != nil {
		fmt.Println("生成验证码时出错:", err)
		//return
	}
	return code
}

// VerifyCode 验证码的有效性
func VerifyCode(secret string, code string) bool {
	//authenticator := googleAuthenticator.GoogleAuthenticator{}
	authenticator := googleAuthenticator.NewGoogleAuthenticator(6)
	// 验证代码的有效性
	isValid := authenticator.VerifyCode(secret, code, 1, 0)
	fmt.Println("验证码是否有效?", isValid)
	return isValid
}

// VerifyCodeMoment 验证是否一致
func VerifyCodeMoment(secret string, code_ string) bool {
	genCode := GetCode(secret)
	fmt.Println("此刻的验证码：", genCode, code_)
	result := strings.Compare(genCode, code_)
	if result == 0 {
		return true
	}
	return false
}

// GenerateQRCode 生成二维码图片，格式为base64
func GenerateQRCode(secret string) string {
	authenticator := googleAuthenticator.NewGoogleAuthenticator(6)
	base64QRCode, err := authenticator.GenerateQRCode("QRCode", secret)
	if err != nil {
		fmt.Println("生成二维码时出错:", err)
		//return
	}
	return base64QRCode
}

// SaveImg 保存二维码图片 主要用于测试
func SaveImg(base64QRCode string) {
	ddd, _ := base64.StdEncoding.DecodeString(base64QRCode) //成图片文件并把文件写入到buffer
	//buffer输出到jpg文件中（不做处理，直接写到文件）
	file, err2 := os.Create("qrcode.png")
	if err2 != nil {
		fmt.Println("保存二维码时出错:", err2)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("保存二维码时出错:", err)
			return
		}
	}(file)
	// 使用Write函数将数据写入文件
	_, err2 = file.Write(ddd)
	if err2 != nil {
		fmt.Println("保存二维码时出错:", err2)
		return
	}
}
