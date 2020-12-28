package repository

import (
	"fmt"
	"github.com/brandon-julio-t/graph-gongular-backend/facades"
	"github.com/brandon-julio-t/graph-gongular-backend/graph/model"
)

type FileUploadRepository struct {
	DB *facades.FileDB
}

func (r *FileUploadRepository) GetById(id string) (*model.FileUpload, error) {
	for _, file := range r.DB.Files {
		if file.ID == id {
			return file, nil
		}
	}

	return nil, fmt.Errorf("file with id %v not found", id)
}

func (r *FileUploadRepository) GetAllByUser(user *model.User) ([]*model.FileUpload, error) {
	files := make([]*model.FileUpload, 0)

	for _, file := range r.DB.Files {
		if file.UserID == user.ID {
			files = append(files, file)
		}
	}

	return files, nil
}

func (r *FileUploadRepository) Save(file *model.FileUpload) (*model.FileUpload, error) {
	r.DB.Files = append(r.DB.Files, file)
	return r.DB.Save(file)
}

func (r *FileUploadRepository) Update(file *model.FileUpload) (*model.FileUpload, error) {
	for i, curr := range r.DB.Files {
		if curr.ID == file.ID {
			r.DB.Files[i] = file
			return file, nil
		}
	}

	return nil, fmt.Errorf("cannot update file %v\n", file)
}

func (r *FileUploadRepository) Delete(id string) (*model.FileUpload, error) {
	var deleted *model.FileUpload = nil

	originalLength := len(r.DB.Files)
	filtered := make([]*model.FileUpload, 0)

	for i := 0; i < originalLength; i++ {
		file := r.DB.Files[i]
		if file.ID != id {
			filtered = append(filtered, file)
		} else {
			deleted = file
		}
	}

	if deleted == nil {
		return nil, fmt.Errorf("cannot delete file with id %v\n", id)
	}

	deleted, err := r.DB.Delete(deleted)
	if err != nil {
		return nil, err
	}

	r.DB.Files = filtered
	return deleted, nil
}
