package services

import (
	"fmt"
	"github.com/brandon-julio-t/graph-gongular-backend/factories/auth-jwt-claims"
	"github.com/brandon-julio-t/graph-gongular-backend/factories/jwt-cookie"
	"github.com/brandon-julio-t/graph-gongular-backend/graph/model"
	"github.com/dgrijalva/jwt-go"
	"net/http"
)

type JwtService struct {
	secret               []byte
	jwtCookieFactory     *jwt_cookie.Factory
	authJwtClaimsFactory *auth_jwt_claims.Factory
}

func NewJwtService(secret []byte) *JwtService {
	return &JwtService{
		secret:               secret,
		jwtCookieFactory:     new(jwt_cookie.Factory),
		authJwtClaimsFactory: new(auth_jwt_claims.Factory),
	}
}

func (s *JwtService) CreateAuthPayload(user *model.User) jwt.Claims {
	return s.authJwtClaimsFactory.NewAuthJwtClaims(user.ID)
}

func (s *JwtService) CreateJwtCookie(token string) *http.Cookie {
	return s.jwtCookieFactory.Create(token)
}

func (s *JwtService) Decode(jwtToken string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return s.secret, nil
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
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS512, payload).SignedString(s.secret)
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

	newData := s.authJwtClaimsFactory.NewAuthJwtClaims(userId)
	return s.Encode(newData)
}

func (s *JwtService) PutJwtCookie(w *http.ResponseWriter, cookie *http.Cookie) {
	(*w).Header().Set("Set-Cookie", cookie.String())
}

func (s *JwtService) ClearJwtCookie(w *http.ResponseWriter) {
	expiredJwtCookie := s.jwtCookieFactory.CreateExpired()
	s.PutJwtCookie(w, expiredJwtCookie)
}
