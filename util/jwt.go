package util

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"hello/setting"
	"strconv"
	"time"
)

type MyClaim struct {
	UserId uint `json:"uid"`
	jwt.StandardClaims
}

func GenToken(userId uint) (string, error) {
	c := MyClaim{
		userId,
		jwt.StandardClaims{
			Id: strconv.Itoa(int(userId)),
			Issuer:    "hello",
			IssuedAt: time.Now().Unix(),
			ExpiresAt: time.Now().Add(setting.TokenExpireDuration).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString(setting.MySecret)
}

func ParseToken(tokenString string) (*MyClaim, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaim{}, func(token *jwt.Token) (interface{}, error) {
		return setting.MySecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaim); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
