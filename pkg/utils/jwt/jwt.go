package jwt

import (
	"errors"
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

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}

func ParseRefreshToken(accessToken, refreshToken string) (newAccessToken, newRefreshToken string, err error) {
	assessClaims, err := ParseToken(accessToken)
	if err != nil {
		return
	}

	refreshClaims, err := ParseToken(refreshToken)
	if err != nil {
		return
	}

	if assessClaims.ExpiresAt > time.Now().Unix() {
		return GenerateToken(assessClaims.ID, assessClaims.UserName)
	}

	if refreshClaims.ExpiresAt > time.Now().Unix() {
		return GenerateToken(assessClaims.ID, assessClaims.UserName)
	}

	return "", "", errors.New("身份过期，请重新登录")
}
