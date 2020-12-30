package services

import (
	"fmt"
	"github.com/brandon-julio-t/graph-gongular-backend/factories/jwt-cookie"
	"github.com/brandon-julio-t/graph-gongular-backend/models"
	"github.com/dgrijalva/jwt-go"
	"net/http"
)

type JwtService struct {
	Secret           []byte
	JwtCookieFactory *jwt_cookie.Factory
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
		return nil, fmt.Errorf("invalid token %v", jwtToken)
	}

	return claims, nil
}

func (s *JwtService) Encode(payload jwt.Claims) (string, error) {
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS512, payload).SignedString(s.Secret)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s *JwtService) Regenerate(oldToken string) (string, error) {
	oldData, err := s.Decode(oldToken)
	if err != nil {
		return "", err
	}

	userId, ok := oldData["userId"].(string)
	if !ok {
		return "", fmt.Errorf("cannot get user id as string %v", oldData)
	}

	newData := models.NewAuthJwtClaims(userId)
	return s.Encode(newData)
}

func (s *JwtService) PutJwtCookie(w *http.ResponseWriter, cookie *http.Cookie) {
	(*w).Header().Set("Set-Cookie", cookie.String())
}

func (s *JwtService) ClearJwtCookie(w *http.ResponseWriter) {
	expiredJwtCookie := s.JwtCookieFactory.CreateExpired()
	s.PutJwtCookie(w, expiredJwtCookie)
}
