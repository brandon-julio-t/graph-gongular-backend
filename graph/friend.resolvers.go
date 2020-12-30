package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/brandon-julio-t/graph-gongular-backend/facades"
	"github.com/brandon-julio-t/graph-gongular-backend/graph/model"
	"github.com/brandon-julio-t/graph-gongular-backend/middlewares"
)

func (r *mutationResolver) AddFriend(ctx context.Context, friendID string) (*model.User, error) {
	if user := middlewares.UseAuth(ctx); user != nil {
		friend, err := r.UserService.GetById(friendID)
		if err != nil {
			return nil, err
		}
		return r.FriendService.Befriend(user, friend)
	}
	return nil, facades.NotAuthenticatedError
}

func (r *mutationResolver) RemoveFriend(ctx context.Context, friendID string) (*model.User, error) {
	if user := middlewares.UseAuth(ctx); user != nil {
		friend, err := r.UserService.GetById(friendID)
		if err != nil {
			return nil, err
		}
		return r.FriendService.Unfriend(user, friend)
	}
	return nil, facades.NotAuthenticatedError
}
