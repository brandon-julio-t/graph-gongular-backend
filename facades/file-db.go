package facades

import (
	"fmt"
	"github.com/brandon-julio-t/graph-gongular-backend/graph/model"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

const BaseDir = "storage"

type FileDB struct {
	Files []*model.FileUpload
}

func (db *FileDB) Save(file *model.FileUpload) (*model.FileUpload, error) {
	if _, err := os.Stat(BaseDir); os.IsNotExist(err) {
		log.Printf("path %v doesn't exists. creating it...", BaseDir)
		if err := os.Mkdir(BaseDir, os.ModePerm); err != nil {
			return nil, err
		}
	}

	data, err := ioutil.ReadAll(file.File)
	if len(data) < 1 {
		return nil, fmt.Errorf("file is empty %v\n", file)
	} else if err != nil {
		return nil, err
	}

	if err := ioutil.WriteFile(file.Path, data, os.ModePerm); err != nil {
		return nil, err
	}

	return file, nil
}

func (db *FileDB) Delete(file *model.FileUpload) (*model.FileUpload, error) {
	filename := fmt.Sprintf("%v.%v", file.ID, file.Extension)
	path := filepath.Join("storage", filename)
	if err := os.Remove(path); err != nil {
		return nil, err
	}

	return file, nil
}


