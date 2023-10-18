package utils

import (
	"time"

	"github.com/baaj2109/webcam_server/config"
	"github.com/baaj2109/webcam_server/global"
	"github.com/golang-jwt/jwt"
)

type Claims struct {
	Email    string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

func GenerateToken(email, password string, cfg *config.JWTConfig) (string, error) {
	now := time.Now()
	expireTime := now.Add(time.Minute * 15).Unix()
	claims := Claims{
		Email:    email,
		Password: global.Md5(password),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime,
			Issuer:    cfg.Issuer,
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(cfg.Secret)
	return token, err
}

func ParseToken(token string, cfg *config.JWTConfig) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return cfg.Secret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
