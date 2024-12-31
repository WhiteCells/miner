package utils

import (
	"crypto/rand"
	"errors"
	"math/big"

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

// +-------------+-----------+------------+
// | field       | key       | value      |
// +-------------+-----------+------------+
// | os:<rig_id> | <farm_id> | <miner_id> |
// +-------------+-----------+------------+

// uid 为生成的矿机索引 ID
// farmID 为矿机所在的矿场 ID
// 生成矿机 ID
// func GenerateRigID(length int, farmID string, minerID string) (string, error) {
// 	const charset = "0123456789"

// 	for {
// 		rigID, err := generateRandomID(charset, length)
// 		if err != nil {
// 			return "", err
// 		}

// 		// 尝试将 ID 存入 Redis，如果已存在则重试
// 		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
// 		defer cancel()

// 		field := fmt.Sprintf("%s:%s", "os", rigID)
// 		success, err := RDB.Client.HSetNX(ctx, field, farmID, minerID).Result()
// 		if err != nil {
// 			return "", err
// 		}

// 		if success {
// 			return rigID, nil
// 		}
// 	}
// }

func GenerateUID() (string, error) {
	node, err := snowflake.NewNode(1) // 1 表示分片标识
	if err != nil {
		return "", err
	}
	id := node.Generate()
	return id.String(), nil
}
