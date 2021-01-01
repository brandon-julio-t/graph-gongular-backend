//go:generate go run github.com/99designs/gqlgen

package graph

import (
	"github.com/brandon-julio-t/graph-gongular-backend/graph/model"
	"github.com/brandon-julio-t/graph-gongular-backend/services"
	"gorm.io/gorm"
	"sync"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	UserService       *services.UserService
	JwtService        *services.JwtService
	FileUploadService *services.FileUploadService
	FriendService     *services.FriendService
	PublicChatService *services.PublicChatService

	Mutex               *sync.Mutex
	PublicChatObservers map[string]chan *model.PublicMessage // user id => public message
}

func NewResolver(db *gorm.DB, secret []byte) *Resolver {
	return &Resolver{
		UserService:       services.NewUserService(db),
		JwtService:        services.NewJwtService(secret),
		FileUploadService: services.NewFileUploadService(db),
		FriendService:     services.NewFriendService(db),
		PublicChatService: services.NewPublicChatService(db),

		Mutex:               new(sync.Mutex),
		PublicChatObservers: make(map[string]chan *model.PublicMessage),
	}
}
