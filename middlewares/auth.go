package middlewares

import (
	"context"
	"fmt"
	"github.com/brandon-julio-t/graph-gongular-backend/graph/model"
	"github.com/brandon-julio-t/graph-gongular-backend/services"
	"github.com/dgrijalva/jwt-go"
	"net/http"
)

type authKeyStruct struct{ name string }

var authKey = authKeyStruct{name: "auth"}

func AuthMiddleware(secret []byte, userService *services.UserService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			user, ok := userByToken(r, secret, userService)

			if !ok {
				next.ServeHTTP(w, r)
				return
			}

			ctx := context.WithValue(r.Context(), authKey, user)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func userByToken(r *http.Request, secret []byte, userService *services.UserService) (*model.User, bool) {
	jwtToken, err := r.Cookie("jwt")

	if err != nil {
		return nil, false
	}

	token, err := jwt.Parse(jwtToken.Value, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return secret, nil
	})

	if err != nil {
		return nil, false
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok || !token.Valid {
		return nil, false
	}

	userId, ok := claims["userId"].(string)

	if !ok {
		return nil, false
	}

	user, err := userService.GetById(userId)

	if err != nil {
		return nil, false
	}

	return user, true
}

func UseAuth(ctx context.Context) *model.User {
	if user, ok := ctx.Value(authKey).(*model.User); ok {
		return user
	}
	return nil
}
