package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/brandon-julio-t/graph-gongular-backend/graph/generated"
	"github.com/brandon-julio-t/graph-gongular-backend/graph/model"
	"github.com/brandon-julio-t/graph-gongular-backend/middlewares"
	jwt "github.com/dgrijalva/jwt-go"
)

func (r *mutationResolver) Register(ctx context.Context, input *model.Register) (*model.User, error) {
	if r.UserService.AlreadyRegistered(input.Email) {
		return nil, errors.New("user already exists")
	}
	return r.UserService.Register(input)
}

func (r *mutationResolver) UpdateAccount(ctx context.Context, input *model.Update) (*model.User, error) {
	return r.UserService.UpdateAccount(input)
}

func (r *mutationResolver) DeleteAccount(ctx context.Context, input *model.DeleteAccount) (*model.User, error) {
	return r.UserService.DeleteAccount(input)
}

func (r *queryResolver) Auth(ctx context.Context) (*model.User, error) {
	if user := middlewares.UseAuth(ctx); user != nil {
		return user, nil
	}
	return nil, errors.New("not signed in")
}

func (r *queryResolver) Login(ctx context.Context, input *model.Login) (*string, error) {
	if user := middlewares.UseAuth(ctx); user != nil {
		return nil, errors.New("already signed in")
	}

	user, err := r.UserService.Login(input.Email, input.Password)

	if err != nil {
		return nil, err
	}

	payload := jwt.MapClaims{"userId": user.ID}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, payload).SignedString(r.JwtSecret)

	if err != nil {
		return nil, err
	}

	cookie := &http.Cookie{
		Name:     "jwt",
		Value:    token,
		Domain:   "graph-gongular-backend.herokuapp.com",
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
	}

	w := *middlewares.UseCookieProvider(ctx)
	w.Header().Set("Set-Cookie", cookie.String())

	return &token, nil
}

func (r *queryResolver) Logout(ctx context.Context) (*bool, error) {
	user := middlewares.UseAuth(ctx)

	if user == nil {
		return nil, errors.New("not signed in")
	}

	cookie := &http.Cookie{
		Name:     "jwt",
		Value:    "",
		Domain:   "graph-gongular-backend.herokuapp.com",
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
		MaxAge:   0,
		Expires:  time.Time{},
	}

	w := *middlewares.UseCookieProvider(ctx)
	w.Header().Add("Set-Cookie", cookie.String())

	result := true
	return &result, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
