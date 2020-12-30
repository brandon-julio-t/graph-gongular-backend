//go:generate go run github.com/99designs/gqlgen

package graph

import (
	"github.com/brandon-julio-t/graph-gongular-backend/services"
	"gorm.io/gorm"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	UserService       *services.UserService
	JwtService        *services.JwtService
	FileUploadService *services.FileUploadService
	FriendService     *services.FriendService
}

func NewResolver(db *gorm.DB, secret []byte) *Resolver {
	return &Resolver{
		UserService: services.NewUserService(db),
		JwtService: services.NewJwtService(secret),
		FileUploadService: services.NewFileUploadService(db),
		FriendService: services.NewFriendService(db),
	}
}
