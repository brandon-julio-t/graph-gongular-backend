package resolver

import (
	fileUpload "github.com/brandon-julio-t/graph-gongular-backend/factories/file-upload"
	jwtCookie "github.com/brandon-julio-t/graph-gongular-backend/factories/jwt-cookie"
	"github.com/brandon-julio-t/graph-gongular-backend/graph"
	"github.com/brandon-julio-t/graph-gongular-backend/repository"
	"github.com/brandon-julio-t/graph-gongular-backend/services"
	"gorm.io/gorm"
)

type Factory struct{}

func (*Factory) Create(secret []byte, db *gorm.DB) *graph.Resolver {
	return &graph.Resolver{
		UserService: &services.UserService{
			UserRepository:     &repository.UserRepository{DB: db},
			UserRoleRepository: &repository.UserRoleRepository{DB: db},
		},
		JwtService: &services.JwtService{
			Secret:           secret,
			JwtCookieFactory: new(jwtCookie.Factory),
		},
		FileUploadService: &services.FileUploadService{
			Factory:    new(fileUpload.Factory),
			Repository: &repository.FileUploadRepository{DB: db},
		},
		FriendService: &services.FriendService{
			Repository: &repository.FriendRepository{DB: db},
		},
	}
}
