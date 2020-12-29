package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/brandon-julio-t/graph-gongular-backend/models"

	"github.com/brandon-julio-t/graph-gongular-backend/facades"
	"github.com/brandon-julio-t/graph-gongular-backend/graph/generated"
	"github.com/brandon-julio-t/graph-gongular-backend/graph/model"
	"github.com/brandon-julio-t/graph-gongular-backend/middlewares"
)

func (r *mutationResolver) Login(ctx context.Context, input *model.Login) (string, error) {
	if user := middlewares.UseAuth(ctx); user != nil {
		return "", facades.AlreadyAuthenticatedError
	}

	user, err := r.UserService.Login(input.Email, input.Password)
	if err != nil {
		return "", err
	}

	payload := models.NewAuthJwtClaims(user.ID)
	token, err := r.JwtService.Encode(payload)
	if err != nil {
		return "", err
	}

	r.JwtService.PutJwtCookie(
		middlewares.UseResponseWriter(ctx),
		r.JwtService.JwtCookieFactory.Create(token),
	)

	return token, nil
}

func (r *mutationResolver) Logout(ctx context.Context) (bool, error) {
	if user := middlewares.UseAuth(ctx); user == nil {
		return false, facades.NotAuthenticatedError
	}

	r.JwtService.ClearJwtCookie(middlewares.UseResponseWriter(ctx))

	return true, nil
}

func (r *queryResolver) Auth(ctx context.Context) (*model.User, error) {
	if user := middlewares.UseAuth(ctx); user != nil {
		return user, nil
	}
	return nil, facades.NotAuthenticatedError
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
