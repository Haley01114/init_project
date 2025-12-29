package utils

import (
	"fmt"
	"time"

	"github.com/Haley01114/init_project/task4/database/models"
	"github.com/golang-jwt/jwt/v5"
)

type JWTConfig struct {
	SecretKey string
	ExpiresIn time.Duration
}

type Claims struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// GenerateToken 生成 JWT token
func GenerateToken(user *models.User) (string, error) {
	cfg := JWTConfig{
		SecretKey: "a-string-secret-at-least-256-bits-long",
		ExpiresIn: 10 * time.Minute,
	}
	claims := Claims{
		UserID:   user.ID,
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   fmt.Sprintf("%d", user.ID),                        //用户标识
			IssuedAt:  jwt.NewNumericDate(time.Now()),                    //签发时间
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(cfg.ExpiresIn)), //过期时间
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(cfg.SecretKey)) //根据密钥进行签名，生成最终 JWT token 字符串
}

// ParseToken 校验token
func ParseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("a-string-secret-at-least-256-bits-long"), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}
