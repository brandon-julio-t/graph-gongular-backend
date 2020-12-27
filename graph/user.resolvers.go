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

func (r *mutationResolver) Register(ctx context.Context, input *model.Register) (*model.User, error) {
	if user := middlewares.UseAuth(ctx); user != nil {
		return nil, errors.New("already signed in")
	}

	if r.UserService.AlreadyRegistered(input.Email) {
		return nil, errors.New("user already exists")
	}

	return r.UserService.Register(input)
}

func (r *mutationResolver) UpdateAccount(ctx context.Context, input *model.Update) (*model.User, error) {
	if user := middlewares.UseAuth(ctx); user != nil {
		return r.UserService.UpdateAccount(user.ID, input)
	}
	return nil, errors.New("not signed in")
}

func (r *mutationResolver) DeleteAccount(ctx context.Context) (*model.User, error) {
	if user := middlewares.UseAuth(ctx); user != nil {
		return r.UserService.DeleteAccount(user.ID)
	}
	return nil, errors.New("not signed in")
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

	token, err := r.JwtService.GenerateAndSetNewTokenInCookie(
		middlewares.UseCookieWriter(ctx),
		user.ID,
	)

	if err != nil {
		return nil, err
	}

	return &token, nil
}

func (r *queryResolver) Logout(ctx context.Context) (*bool, error) {
	if user := middlewares.UseAuth(ctx); user == nil {
		return nil, errors.New("not signed in")
	}

	r.JwtService.SetExpiredTokenInCookie(middlewares.UseCookieWriter(ctx))

	result := true
	return &result, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
