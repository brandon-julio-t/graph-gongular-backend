package repository

import (
	"github.com/brandon-julio-t/graph-gongular-backend/graph/model"
	"gorm.io/gorm"
)

type FileUploadRepository struct {
	DB *gorm.DB
}

func (r *FileUploadRepository) GetById(id string) (*model.FileUpload, error) {
	fileUpload := new(model.FileUpload)
	if err := r.DB.Preload("User").First(fileUpload, "file_uploads.id = ?", id).Error; err != nil {
		return nil, err
	}
	return fileUpload, nil
}

func (r *FileUploadRepository) Save(fileUpload *model.FileUpload) (*model.FileUpload, error) {
	if err := r.DB.Create(fileUpload).Error; err != nil {
		return nil, err
	}
	return fileUpload, nil
}

func (r *FileUploadRepository) Update(fileUpload *model.FileUpload) (*model.FileUpload, error) {
	if err := r.DB.Save(fileUpload).Error; err != nil {
		return nil, err
	}
	return fileUpload, nil
}

func (r *FileUploadRepository) Delete(fileUpload *model.FileUpload) (*model.FileUpload, error) {
	if err := r.DB.Delete(fileUpload).Error; err != nil {
		return nil, err
	}
	return fileUpload, nil
}
