package file_upload

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/brandon-julio-t/graph-gongular-backend/graph/model"
	"github.com/google/uuid"
	"io/ioutil"
	"strings"
)

type Factory struct{}

func (f *Factory) Create(file *graphql.Upload, user *model.User) (*model.FileUpload, error) {
	filenameSplit := strings.Split(file.Filename, ".")
	filename := strings.Join(filenameSplit[:len(filenameSplit)-1], "")
	extension := filenameSplit[len(filenameSplit)-1]

	id := uuid.Must(uuid.NewRandom()).String()

	data, err := ioutil.ReadAll(file.File)
	if err != nil {
		return nil, err
	}

	fileUpload := &model.FileUpload{
		ID:          id,
		File:        data,
		Filename:    filename,
		Size:        file.Size,
		ContentType: file.ContentType,
		Extension:   extension,
		UserID:      user.ID,
		User:        user,
	}

	fileUpload.Filename = filename
	return fileUpload, nil
}
