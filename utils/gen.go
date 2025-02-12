package utils

import (
	"crypto/rand"
	"crypto/sha1"
	"encoding/hex"
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

func GenerateUID() (string, error) {
	node, err := snowflake.NewNode(1) // 1 表示分片标识
	if err != nil {
		return "", err
	}
	id := node.Generate() // 不会以 0 开头
	return id.String(), nil
}

// 28a0259f06d50b134d1c90bd11521ceb0d9fc282
// 3c658db850c04e9728e0ff51f5116998af8e1ae5
func GenerateFarmHash(data string) string {
	hash := sha1.New()
	hash.Write([]byte(data))
	rawHash := hash.Sum(nil)
	return hex.EncodeToString(rawHash)
}
