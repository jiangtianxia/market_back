package utils

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	myKey     []byte
	jwtExpire int
	issuer    string
)

func InitJwt(key string, expire int, iss string) {
	myKey = []byte(key)
	jwtExpire = expire
	issuer = iss
}

func GetIssuer() string {
	return issuer
}

type UserClaims struct {
	Openid string `json:"openid"`
	jwt.StandardClaims
}

// 生成token
func GenerateToken(openid string) (string, error) {
	userClaim := &UserClaims{
		Openid: openid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(
				time.Duration(jwtExpire) * time.Hour).Unix(),
			Issuer: issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaim)
	tokenString, err := token.SignedString(myKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// 解析token
func AnalyzeToken(tokenString string) (*UserClaims, error) {
	UserClaims := new(UserClaims)

	claims, err := jwt.ParseWithClaims(tokenString, UserClaims, func(token *jwt.Token) (interface{}, error) {
		return myKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !claims.Valid {
		return nil, fmt.Errorf("analyze Token Error:%v", err)
	}
	return UserClaims, nil
}
