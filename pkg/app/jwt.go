package app

import (
	"time"

	"github.com/dgrijalva/jwt-go"

	"blog/global"
	"blog/pkg/util"
)

type Claims struct {
	AppId     string `json:"app_id"`
	AppSecret string `json:"app_secret"`
	jwt.StandardClaims
}

func GetJWTSecret() []byte {
	return []byte(global.JwtSetting.Secret)
}

func GenerateToken(appId, appSecret string) (string, error) {
	newTime := time.Now()
	expireTime := newTime.Add(global.JwtSetting.Expire)
	claims := Claims{
		AppId:     appId,
		AppSecret: util.EnCodeMd5(appSecret),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    global.JwtSetting.Issuer,
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return tokenClaims.SignedString(GetJWTSecret())
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return GetJWTSecret(), nil
	})

	if tokenClaims != nil {
		claims, ok := tokenClaims.Claims.(*Claims)
		if ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
