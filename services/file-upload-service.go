package services

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/brandon-julio-t/graph-gongular-backend/factories/file-upload"
	"github.com/brandon-julio-t/graph-gongular-backend/graph/model"
	"github.com/brandon-julio-t/graph-gongular-backend/repository"
	"gorm.io/gorm"
)

type FileUploadService struct {
	Factory    *file_upload.Factory
	Repository *repository.FileUploadRepository
}

func NewFileUploadService(db *gorm.DB) *FileUploadService {
	return &FileUploadService{
		Factory:    new(file_upload.Factory),
		Repository: &repository.FileUploadRepository{DB: db},
	}
}

func (s *FileUploadService) GetFileById(id string) (*model.FileUpload, error) {
	return s.Repository.GetById(id)
}

func (s *FileUploadService) SaveFile(file *graphql.Upload, user *model.User) (*model.FileUpload, error) {
	fileUpload, err := s.Factory.Create(file, user)
	if err != nil {
		return nil, err
	}
	return s.Repository.Save(fileUpload)
}

func (s *FileUploadService) UpdateFile(input *model.UpdateFile) (*model.FileUpload, error) {
	file, err := s.Repository.GetById(input.ID)
	if err != nil {
		return nil, err
	}

	file.Filename = input.Filename

	return s.Repository.Update(file)
}

func (s *FileUploadService) DeleteFile(id string) (*model.FileUpload, error) {
	fileUpload, err := s.Repository.GetById(id)
	if err != nil {
		return nil, err
	}
	return s.Repository.Delete(fileUpload)
}
