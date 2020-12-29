package middlewares

import (
	"context"
	"fmt"
	"github.com/brandon-julio-t/graph-gongular-backend/factories/jwt-cookie"
	"github.com/brandon-julio-t/graph-gongular-backend/graph/model"
	"github.com/brandon-julio-t/graph-gongular-backend/services"
	"log"
	"net/http"
)

type authKeyStruct struct{ name string }

var authKey = authKeyStruct{name: "auth"}

func AuthMiddleware(jwtService *services.JwtService, userService *services.UserService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			user, err := marshalTokenIntoUser(r, jwtService, userService)

			if err != nil {
				log.Println(err)
				next.ServeHTTP(w, r)
				return
			}

			ctx := context.WithValue(r.Context(), authKey, user)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func marshalTokenIntoUser(
	r *http.Request,
	jwtService *services.JwtService,
	userService *services.UserService,
) (*model.User, error) {
	jwtToken, err := r.Cookie(jwt_cookie.JwtCookieName)
	if err != nil {
		return nil, err
	}

	payload, err := jwtService.Decode(jwtToken.Value)
	if err != nil {
		return nil, err
	}

	userId, ok := payload["userId"].(string)
	if !ok {
		return nil, fmt.Errorf("cannot find userId in token %v\n", payload)
	}

	return userService.GetById(userId)
}

func UseAuth(ctx context.Context) *model.User {
	if user, ok := ctx.Value(authKey).(*model.User); ok {
		return user
	}
	return nil
}
