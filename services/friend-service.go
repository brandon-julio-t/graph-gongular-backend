package services

import (
	"github.com/brandon-julio-t/graph-gongular-backend/graph/model"
	"github.com/brandon-julio-t/graph-gongular-backend/repository"
)

type FriendService struct {
	Repository *repository.FriendRepository
}

func (s *FriendService) Befriend(user, friend *model.User) (*model.User, error) {
	return s.Repository.Save(user, friend)
}

func (s *FriendService) Unfriend(user, friend *model.User) (*model.User, error) {
	return s.Repository.Delete(user, friend)
}

