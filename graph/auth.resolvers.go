package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"

	"github.com/brandon-julio-t/graph-gongular-backend/graph/generated"
	"github.com/brandon-julio-t/graph-gongular-backend/graph/model"
	"github.com/brandon-julio-t/graph-gongular-backend/middlewares"
)

func (r *mutationResolver) Login(ctx context.Context, input *model.Login) (string, error) {
	if user := middlewares.UseAuth(ctx); user != nil {
		return "", errors.New("already signed in")
	}

	user, err := r.UserService.Login(input.Email, input.Password)
	if err != nil {
		return "", err
	}

	token, err := r.JwtService.GenerateAndSetNewTokenInCookie(
		middlewares.UseCookieWriter(ctx),
		user.ID,
	)

	if err != nil {
		return "", err
	}

	return token, nil
}

func (r *mutationResolver) Logout(ctx context.Context) (bool, error) {
	if user := middlewares.UseAuth(ctx); user == nil {
		return false, errors.New("not authenticated")
	}

	r.JwtService.SetExpiredTokenInCookie(middlewares.UseCookieWriter(ctx))

	return true, nil
}

func (r *queryResolver) Auth(ctx context.Context) (*model.User, error) {
	if user := middlewares.UseAuth(ctx); user != nil {
		return user, nil
	}
	return nil, errors.New("not authenticated")
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
