package jwt

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

/*
	-------------use-------------------

	token, _ := jwt.GenerateToken(12312)
	res, err := jwt.ParseToken(token)
	if err != nil {
		fmt.Println("解析失败")
	} else {
		fmt.Println(res.Uid)
	}
	------------------------------------
*/

// 要求
type Claim struct {
	Uid int64
	jwt.StandardClaims
}

type JWT struct {
	Secret string `yaml:"secret,omitempty"`
	Expire int    `yaml:"expire,omitempty"`
}

func GenerateToken(uid int64) (string, error) {
	// 过期时间60秒
	expire := time.Now().Add(time.Duration(60) * time.Second)
	claims := Claim{
		Uid: uid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expire.Unix(),
			Issuer: "Will_Test",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	fmt.Println(tokenClaims)
	var jwt JWT
	fmt.Println(tokenClaims.SigningString())
	fmt.Println(tokenClaims.SignedString([]byte(jwt.Secret)))
	return tokenClaims.SignedString([]byte(jwt.Secret))
}

func ParseToken(token string) (claims *Claim, err error) {
	var jwtStruct JWT
	tokenClaims, err := jwt.ParseWithClaims(
		token,
		&Claim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtStruct.Secret), nil
		},
	)
	if err != nil {
		return
	}
	if !tokenClaims.Valid {
		err = errors.New("token 无效")
		return
	}
	claims, ok := tokenClaims.Claims.(*Claim)
	if !ok {
		err = errors.New("token 解析失败")
	}
	return
}