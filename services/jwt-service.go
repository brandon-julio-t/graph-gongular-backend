package services

import (
	"fmt"
	"github.com/brandon-julio-t/graph-gongular-backend/factories"
	"github.com/brandon-julio-t/graph-gongular-backend/models"
	"github.com/dgrijalva/jwt-go"
	"net/http"
)

type JwtService struct {
	Secret           []byte
	JwtCookieFactory *factories.JwtCookieFactory
}

func (s *JwtService) Decode(jwtToken string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return s.Secret, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token %v\n", jwtToken)
	}

	return claims, nil
}

func (s *JwtService) GenerateAndSetNewTokenInCookie(w *http.ResponseWriter, userId string) (string, error) {
	token, err := s.encode(models.NewAuthJwtClaims(userId))
	if err != nil {
		return "", err
	}

	s.setTokenInCookie(w, s.JwtCookieFactory.NewJwtCookie(token))
	return token, nil
}

func (s *JwtService) encode(payload jwt.Claims) (string, error) {
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS512, payload).SignedString(s.Secret)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s *JwtService) SetExpiredTokenInCookie(w *http.ResponseWriter) {
	s.setTokenInCookie(w, s.JwtCookieFactory.NewExpiredJwtCookie())
}

func (s *JwtService) setTokenInCookie(w *http.ResponseWriter, cookie *http.Cookie) {
	(*w).Header().Set("Set-Cookie", cookie.String())
}
