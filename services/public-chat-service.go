package services

import (
	"github.com/brandon-julio-t/graph-gongular-backend/graph/model"
	"github.com/brandon-julio-t/graph-gongular-backend/repository"
	"gorm.io/gorm"
)

type PublicChatService struct {
	Repository *repository.PublicChatRepository
}

func NewPublicChatService(db *gorm.DB) *PublicChatService {
	return &PublicChatService{
		Repository: &repository.PublicChatRepository{DB: db},
	}
}

func (s *PublicChatService) All() ([]*model.PublicMessage, error) {
	return s.Repository.GetAll()
}

func (s *PublicChatService) Save(message *model.PublicMessage) (*model.PublicMessage, error) {
	return s.Repository.Save(message)
}
