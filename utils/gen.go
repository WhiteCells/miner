package utils

import (
	"crypto/rand"
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"math/big"

	"github.com/bwmarrin/snowflake"
	"github.com/google/uuid"
)

func GenerateRigPass(length int) (string, error) {
	if length < 8 {
		return "", errors.New("invalid argument")
	}
	const charset = "abcdefghjkmnpqrstuvwxyzABCDEFGHJKMNPQRSTUVWXYZ23456789"
	str := make([]byte, length)
	for i := range str {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", errors.New("rand int error")
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
func GenerateFarmHash() string {
	hash := sha1.New()
	hash.Write([]byte(uuid.New().String()))
	return hex.EncodeToString(hash.Sum(nil))
}
