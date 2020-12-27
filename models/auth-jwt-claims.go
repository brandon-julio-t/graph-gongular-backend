package models

import (
	"github.com/dgrijalva/jwt-go"
	"log"
	"time"
)

type AuthJwtClaims struct {
	*jwt.StandardClaims
	UserId string `json:"userId"`
}

func NewAuthJwtClaims(userId string) *AuthJwtClaims {
	duration, err := time.ParseDuration("15m")

	if err != nil  {
		log.Fatalln(err)
	}

	return &AuthJwtClaims{
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: time.Now().Add(duration).Unix(),
			IssuedAt:  time.Now().Unix(),
			NotBefore: time.Now().Unix(),
		},
		UserId: userId,
	}
}
