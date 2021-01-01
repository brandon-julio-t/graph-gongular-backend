package auth_jwt_claims

import (
	"github.com/dgrijalva/jwt-go"
	"log"
	"time"
)

type Factory struct{}

type authJwtClaims struct {
	*jwt.StandardClaims
	UserId string `json:"userId"`
}

func (*Factory) NewAuthJwtClaims(userId string) jwt.Claims {
	duration, err := time.ParseDuration("15m")

	if err != nil {
		log.Fatalln(err)
	}

	return &authJwtClaims{
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: time.Now().Add(duration).Unix(),
			IssuedAt:  time.Now().Unix(),
			NotBefore: time.Now().Unix(),
		},
		UserId: userId,
	}
}
