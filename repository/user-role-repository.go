package repository

import (
	"github.com/brandon-julio-t/graph-gongular-backend/graph/model"
	"gorm.io/gorm"
)

type UserRoleRepository struct {
	DB *gorm.DB
}

func (r *UserRoleRepository) GetUserRole() (*model.UserRole, error) {
	role := &model.UserRole{}
	if err := r.DB.First(role, "name = ?", "User").Error; err != nil {
		return nil, err
	}
	return role, nil
}
