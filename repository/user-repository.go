package repository

import (
	"github.com/brandon-julio-t/graph-gongular-backend/facades"
	"github.com/brandon-julio-t/graph-gongular-backend/graph/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
	FileRepository *FileRepository
	FileDB *facades.FileDB
}

func (r *UserRepository) GetById(id string) (*model.User, error) {
	user := new(model.User)

	if err := r.DB.Joins("UserRole").First(user, "users.id = ?", id).Error; err != nil {
		return nil, err
	}

	files, err := r.FileRepository.GetByUser(user)
	if err != nil {
		return nil, err
	}

	user.Files = files

	return user, nil
}

func (r *UserRepository) GetByEmail(email string) (*model.User, error) {
	user := new(model.User)

	if err := r.DB.Joins("UserRole").First(user, "email = ?", email).Error; err != nil {
		return nil, err
	}

	files, err := r.FileRepository.GetByUser(user)
	if err != nil {
		return nil, err
	}

	user.Files = files

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
	if err := r.DB.Where("id = ?", user.ID).Delete(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
