package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"time"

	"github.com/brandon-julio-t/graph-gongular-backend/facades"
	"github.com/brandon-julio-t/graph-gongular-backend/graph/generated"
	"github.com/brandon-julio-t/graph-gongular-backend/graph/model"
	"github.com/brandon-julio-t/graph-gongular-backend/middlewares"
	"github.com/google/uuid"
)

func (r *mutationResolver) SendMessage(ctx context.Context, message string) (*model.PublicMessage, error) {
	if user := middlewares.UseAuth(ctx); user != nil {
		publicMessage, err := r.PublicChatService.Save(&model.PublicMessage{
			ID:        uuid.Must(uuid.NewRandom()).String(),
			UserID:    user.ID,
			User:      user,
			Message:   message,
			CreatedAt: time.Now(),
		})

		if err != nil {
			return nil, err
		}

		r.Mutex.Lock()
		for _, observer := range r.PublicChatObservers {
			observer <- publicMessage
		}
		r.Mutex.Unlock()

		return publicMessage, nil
	}

	return nil, facades.NotAuthenticatedError
}

func (r *queryResolver) Messages(ctx context.Context) ([]*model.PublicMessage, error) {
	if user := middlewares.UseAuth(ctx); user == nil {
		return nil, facades.NotAuthenticatedError
	}

	r.Mutex.Lock()
	all, err := r.PublicChatService.All()
	r.Mutex.Unlock()

	return all, err
}

func (r *subscriptionResolver) MessageAdded(ctx context.Context) (<-chan *model.PublicMessage, error) {
	if user := middlewares.UseAuth(ctx); user != nil {
		observer := make(chan *model.PublicMessage, 1)

		r.Mutex.Lock()
		r.PublicChatObservers[user.ID] = observer
		r.Mutex.Unlock()

		go func() {
			<-ctx.Done()
			r.Mutex.Lock()
			delete(r.PublicChatObservers, user.ID)
			r.Mutex.Unlock()
		}()

		return observer, nil
	}

	return nil, facades.NotAuthenticatedError
}

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

type subscriptionResolver struct{ *Resolver }
