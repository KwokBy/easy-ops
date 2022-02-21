// Package jwt jwt 认证包
package jwt

import (
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Data struct {
	UserID int64 `json:"user_id"`
}

// Token JWT
type Token struct {
	*jwt.RegisteredClaims
	Data
}

// New 创建一个新的Token
func New() string {
	// 使用SigningMethodHS256生成签名的方法
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, Token{
		&jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
		Data{
			1,
		},
	})
	// TODO 签名用配置文件管理
	token, err := claims.SignedString([]byte("secret"))
	if err != nil {
		panic(err)
	}
	return token
}

// IsValid 校验token是否有效
func IsValid(token string) bool {
	tt, err := jwt.ParseWithClaims(token, &Token{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// TODO 签名用配置文件管理
		return []byte("secret"), nil
	})
	if err != nil {
		fmt.Println(err)
		return false
	}
	if !tt.Valid {
		return false
	}
	return true
}

// GetUserIDFromToken 获取用户ID
func GetUserIDFromToken(token string) (int64, error) {
	// TODO 签名用配置文件管理
	tt, err := jwt.ParseWithClaims(token, &Token{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("secret"), nil
	})
	if err != nil {
		return -1, fmt.Errorf("parse jwt failed: %v", err)
	}

	if claims, ok := tt.Claims.(*Token); ok && tt.Valid {
		return claims.Data.UserID, nil
	}
	return -1, fmt.Errorf("failed to get userid")
}

// GetUserIDFromHTTPRequest 从HTTP请求中的jwt获取UserID
func GetUserIDFromHTTPRequest(r *http.Request) (int64, error) {
	token := r.Header.Get("Authorization")
	return GetUserIDFromToken(token)
}
