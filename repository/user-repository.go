package repository

import (
	"github.com/brandon-julio-t/graph-gongular-backend/graph/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func (r *UserRepository) GetAllExcept(user *model.User) ([]*model.User, error) {
	var users []*model.User
	if err := r.preloadUserAssociations().Where("users.id != ?", user.ID).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserRepository) preloadUserAssociations() *gorm.DB {
	return r.DB.
		Preload("UserRole").
		Preload("Friends").
		Preload("Friends.UserRole").
		Preload("FileUploads")
}

func (r *UserRepository) GetById(id string) (*model.User, error) {
	user := new(model.User)
	if err := r.preloadUserAssociations().First(user, "users.id = ?", id).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) GetByEmail(email string) (*model.User, error) {
	user := new(model.User)
	if err := r.preloadUserAssociations().First(user, "email = ?", email).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) Save(user *model.User) (*model.User, error) {
	if err := r.DB.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) Update(user *model.User) (*model.User, error) {
	if err := r.DB.Save(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) Delete(user *model.User) (*model.User, error) {
	if err := r.DB.Delete(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
