package main

import (
	"crypto/rand"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type claimType struct {
	User
	jwt.StandardClaims
}

type JWTHandler struct {
	signingKey []byte
}

func NewJWT() JWTHandler {
	token := make([]byte, 16)
	rand.Read(token)
	return JWTHandler{token}
}

func (jh JWTHandler) GenerateJWT(st User) (string, error) {
	claims := claimType{
		st,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jh.signingKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (jh JWTHandler) verifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(tokenString, &claimType{}, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jh.signingKey, nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func (jh JWTHandler) DecodeJWT(tokenString string) (st User, err error) {
	token, err := jh.verifyToken(tokenString)
	if err != nil {
		err = errors.New("invalid token")
		return
	}
	claims, ok := token.Claims.(*claimType)
	if ok && token.Valid {
		if time.Now().Unix() > claims.ExpiresAt {
			err = errors.New("token expired")
			return
		}

		st = claims.User

		return
	}
	return
}
