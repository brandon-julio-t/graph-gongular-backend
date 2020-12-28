package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"github.com/brandon-julio-t/graph-gongular-backend/facades"

	"github.com/brandon-julio-t/graph-gongular-backend/graph/generated"
	"github.com/brandon-julio-t/graph-gongular-backend/graph/model"
	"github.com/brandon-julio-t/graph-gongular-backend/middlewares"
)

func (r *mutationResolver) Register(ctx context.Context, input *model.Register) (*model.User, error) {
	if user := middlewares.UseAuth(ctx); user != nil {
		return nil, facades.AlreadyAuthenticatedError
	}

	if r.UserService.AlreadyRegistered(input.Email) {
		return nil, errors.New("user already exists")
	}

	return r.UserService.Register(input)
}

func (r *mutationResolver) UpdateAccount(ctx context.Context, input *model.UpdateUser) (*model.User, error) {
	if user := middlewares.UseAuth(ctx); user != nil {
		return r.UserService.UpdateAccount(user.ID, input)
	}
	return nil, facades.NotAuthenticatedError
}

func (r *mutationResolver) DeleteAccount(ctx context.Context) (*model.User, error) {
	if user := middlewares.UseAuth(ctx); user != nil {
		return r.UserService.DeleteAccount(user.ID)
	}
	return nil, facades.NotAuthenticatedError
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
