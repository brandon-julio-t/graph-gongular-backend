package repository

import (
	"fmt"
	"github.com/brandon-julio-t/graph-gongular-backend/facades"
	"github.com/brandon-julio-t/graph-gongular-backend/graph/model"
)

type FileRepository struct {
	DB *facades.FileDB
}

func (r *FileRepository) GetById(id string) (*model.FileUpload, error) {
	for _, file := range r.DB.Files {
		if file.ID == id {
			return file, nil
		}
	}

	return nil, fmt.Errorf("file with id %v not found", id)
}

func (r *FileRepository) GetByUser(user *model.User) ([]*model.FileUpload, error) {
	files := make([]*model.FileUpload, 0)

	for _, file := range r.DB.Files {
		if file.UserID == user.ID {
			files = append(files, file)
		}
	}

	return files, nil
}

func (r *FileRepository) Save(file *model.FileUpload) (*model.FileUpload, error) {
	return r.DB.Save(file)
}

func (r *FileRepository) Update(file *model.FileUpload) (*model.FileUpload, error) {
	for i, curr := range r.DB.Files {
		if curr.ID == file.ID {
			r.DB.Files[i] = file
			return file, nil
		}
	}

	return nil, fmt.Errorf("cannot update file %v\n", file)
}

func (r *FileRepository) Delete(id string) (*model.FileUpload, error) {
	var deleted *model.FileUpload = nil

	originalLength := len(r.DB.Files)
	filtered := make([]*model.FileUpload, originalLength-1)
	curr := 0

	for i := 0; i < originalLength; i++ {
		file := r.DB.Files[i]
		if file.ID != id {
			filtered[curr] = file
			curr++
		} else {
			deleted = file
		}
	}

	if deleted == nil {
		return nil, fmt.Errorf("cannot delete file with id %v\n", id)
	}

	r.DB.Files = filtered

	return deleted, nil
}
