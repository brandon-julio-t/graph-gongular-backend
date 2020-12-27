package facades

import "github.com/brandon-julio-t/graph-gongular-backend/graph/model"

type FileDB struct {
	Files []*model.FileUpload
}

func (db *FileDB) Save(file *model.FileUpload) (*model.FileUpload, error) {
	db.Files = append(db.Files, file)
	return file, nil
}
