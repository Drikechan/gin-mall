package jwt

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"test-gin-mall/consts"
	"time"
)

var jwtSecret = []byte("FanOne")

type Claims struct {
	ID       uint   `json:"id"`
	UserName string `json:"userName"`
	jwt.StandardClaims
}

func GenerateToken(id uint, name string) (accessToken, refreshToken string, err error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(consts.AccessTokenExpireDuration)
	rtExpireTime := nowTime.Add(consts.AccessRefreshTokenExpireDuration)
	claims := Claims{
		ID:       id,
		UserName: name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "mall",
		},
	}
	accessToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(jwtSecret)
	fmt.Println(accessToken)
	if err != nil {
		return "", "", err
	}

	refreshToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: rtExpireTime.Unix(),
		Issuer:    "mall",
	}).SignedString(jwtSecret)

	if err != nil {
		return "", "", err
	}
	return accessToken, refreshToken, err
}
