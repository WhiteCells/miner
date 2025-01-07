package utils

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
	mr "math/rand"
	"time"

	"github.com/bwmarrin/snowflake"
)

func GeneratePass(length int) (string, error) {
	if length < 8 {
		return "", errors.New("invalid argument")
	}
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	str := make([]byte, length)
	for i := range str {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		str[i] = charset[num.Int64()]
	}
	return string(str), nil
}

func GenerateUID() (string, error) {
	node, err := snowflake.NewNode(1) // 1 表示分片标识
	if err != nil {
		return "", err
	}
	id := node.Generate() // 不会以 0 开头
	return id.String(), nil
}

func GenerateTaskID() string {
	// 获取当前时间戳后8位
	timestamp := time.Now().UnixNano() / 1e6 % 100000000

	// 生成4位随机数
	mr.Seed(time.Now().UnixNano())
	randomNum := mr.Intn(9000) + 1000 // 生成1000-9999之间的随机数

	// 组合成12位数的ID
	return fmt.Sprintf("%d%d", timestamp, randomNum)
}
