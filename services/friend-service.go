package services

import (
	"github.com/brandon-julio-t/graph-gongular-backend/graph/model"
	"github.com/brandon-julio-t/graph-gongular-backend/repository"
	"gorm.io/gorm"
)

type FriendService struct {
	repository *repository.FriendRepository
}

func NewFriendService(db *gorm.DB) *FriendService {
	return &FriendService{
		repository: &repository.FriendRepository{DB: db},
	}
}

func (s *FriendService) Befriend(user, friend *model.User) (*model.User, error) {
	return s.repository.Save(user, friend)
}

func (s *FriendService) Unfriend(user, friend *model.User) (*model.User, error) {
	return s.repository.Delete(user, friend)
}
