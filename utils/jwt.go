package utils

import (
	"errors"
	"miner/common/role"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	UserID   string        `json:"user_id"`
	Username string        `json:"user_name"`
	Role     role.RoleType `json:"user_role"`
	jwt.RegisteredClaims
}

var jwtSecret []byte

func InitJWT() {
	jwtSecret = []byte(Config.JWT.Secret)
}

// 生成 JWT token
// token 无状态
// todo 防止删除的用户生成的 token 被应用于新用户，需要在表中进行确认
// todo 表中确认效率太低了，维护一个黑名单，存放过期 token
func GenerateToken(userID string, username string, role role.RoleType, expireHours int) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(time.Duration(expireHours) * time.Hour)

	claims := Claims{
		UserID:   userID,
		Username: username,
		Role:     role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireTime),
			IssuedAt:  jwt.NewNumericDate(nowTime),
			Issuer:    "miner",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

// 解析 JWT token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
