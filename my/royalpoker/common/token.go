package common

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

const secret = "Y/IeiwJdSZd+WDPB73J5zOfqR9/tLcehDLTrbv96SOGIcag9YJS96r6mMwE5ixM4bWLH5fh5jYKg0L1Iy1Vrcg=="

type Claims struct {
	jwt.StandardClaims `json:"standard_claims"`
	Uid                int    `json:"uid"`
	Name               string `json:"name"`
	IsAdmin            bool   `json:"is_admin"`
}

func Encode(claims *Claims) (string, error) {
	claims.StandardClaims.ExpiresAt = time.Now().Add(24 * time.Hour).Unix()
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), claims)
	return token.SignedString([]byte(secret))
}

func Decode(token string) (*Claims, error) {
	var err error
	var claims = &Claims{}
	_, err = jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	return claims, err
}
