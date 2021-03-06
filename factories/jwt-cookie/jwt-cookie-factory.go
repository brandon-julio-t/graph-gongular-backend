package jwt_cookie

import (
	"net/http"
	"os"
	"time"
)

const (
	JwtCookieName = "jwt"
	domain        = "graph-gongular-backend.herokuapp.com"
)

type Factory struct{}

func (*Factory) Create(token string) *http.Cookie {
	if os.Getenv("APP_ENV") == "development" {
		return &http.Cookie{
			Name:     JwtCookieName,
			Value:    token,
			HttpOnly: true,
		}
	}

	return &http.Cookie{
		Name:     JwtCookieName,
		Value:    token,
		Domain:   domain,
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
	}
}

func (*Factory) CreateExpired() *http.Cookie {
	if os.Getenv("APP_ENV") == "development" {
		return &http.Cookie{
			Name:     JwtCookieName,
			Value:    "",
			HttpOnly: true,
			MaxAge:   0,
			Expires:  time.Time{},
		}
	}

	return &http.Cookie{
		Name:     JwtCookieName,
		Value:    "",
		Domain:   domain,
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
		MaxAge:   0,
		Expires:  time.Time{},
	}
}
