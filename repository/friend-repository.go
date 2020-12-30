package repository

import (
	"github.com/brandon-julio-t/graph-gongular-backend/graph/model"
	"gorm.io/gorm"
)

type FriendRepository struct {
	DB *gorm.DB
}

func (r FriendRepository) Save(user, friend *model.User) (*model.User, error) {
	if err := r.DB.Model(user).Association("Friends").Append(friend); err != nil {
		return nil, err
	}
	return friend, nil
}

func (r FriendRepository) Delete(user, friend *model.User) (*model.User, error) {
	if err := r.DB.Model(user).Association("Friends").Delete(friend); err != nil {
		return nil, err
	}
	return friend, nil
}
