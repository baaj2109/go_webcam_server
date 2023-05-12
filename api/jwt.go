package api

import (
	"time"

	"github.com/baaj2109/webcam_server/global"
	"github.com/baaj2109/webcam_server/settings"
	"github.com/golang-jwt/jwt"
)

type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

func GenerateToken(username, password string, cfg *settings.JWTConfig) (string, error) {
	now := time.Now()
	expireTime := now.Add(time.Minute * 15).Unix()
	claims := Claims{
		Username: global.Md5(username),
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

func ParseToken(token string, cfg *settings.JWTConfig) (*Claims, error) {
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
