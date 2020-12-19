package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

//type Credentials struct {
//	Password string `json:"password"`
//	Username string `json:"username"`
//}

var jwtKey = []byte("my_secret_key")

type Claims struct {
	Uuid string `json:"uuid"`
	Exp  int64  `json:"exp"`
	jwt.StandardClaims
}

func Connecings(uuid, login, pwd string) (tokenString string, expirationTime time.Time, err error) {
	if login != "css" && pwd != "Mince1234" {
		return
	}
	expirationTime = time.Now().Add(1440 * time.Minute)
	ee := expirationTime.Unix()
	claims := &Claims{
		Uuid: uuid,
		Exp:  ee,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: ee,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(jwtKey)
	return
}

func GetJwt(tknStr string) (claims *Claims, err error) {
	claims = &Claims{}
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			err = fmt.Errorf("unautorized")
		}
		return
	}
	if !tkn.Valid {
		err = fmt.Errorf("unautorized")
		return
	}
	return
}

func Refresh(tknStr string) (tokenString string, expirationTime time.Time, err error) {
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			err = fmt.Errorf("unautorized")
		}
		return
	}
	if !tkn.Valid {
		err = fmt.Errorf("unautorized")
	}
	if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 30*time.Second {
		err = fmt.Errorf("bad request")
		return
	}
	expirationTime = time.Now().Add(1440 * time.Minute)
	claims.ExpiresAt = expirationTime.Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(jwtKey)
	return
}
