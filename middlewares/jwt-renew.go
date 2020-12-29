package middlewares

import (
	jwtCookie "github.com/brandon-julio-t/graph-gongular-backend/factories/jwt-cookie"
	"github.com/brandon-julio-t/graph-gongular-backend/services"
	"net/http"
)

func JwtRenewMiddleware(jwtService *services.JwtService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			renewJwtToken(w, r, jwtService)
			next.ServeHTTP(w, r)
		})
	}
}

func renewJwtToken(w http.ResponseWriter, r *http.Request, jwtService *services.JwtService) {
	oldJwtCookie, err := r.Cookie(jwtCookie.JwtCookieName)
	if err != nil {
		return
	}

	oldToken := oldJwtCookie.Value
	newToken, err := jwtService.Regenerate(oldToken)
	if err != nil {
		return
	}

	newJwtCookie := new(jwtCookie.Factory).Create(newToken)
	w.Header().Set("Set-Cookie", newJwtCookie.String())
}
