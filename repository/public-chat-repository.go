package repository

import (
	"github.com/brandon-julio-t/graph-gongular-backend/graph/model"
	"gorm.io/gorm"
)

type PublicChatRepository struct {
	DB *gorm.DB
}

func (r *PublicChatRepository) GetAll() ([]*model.PublicMessage, error) {
	var messages []*model.PublicMessage
	if err := r.DB.Preload("User").Preload("User.UserRole").Find(&messages).Error; err != nil {
		return nil, err
	}
	return messages, nil
}

func (r *PublicChatRepository) Save(message *model.PublicMessage) (*model.PublicMessage, error) {
	if err := r.DB.Create(message).Error; err != nil {
		return nil, err
	}
	return message, nil
}
