package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/99designs/gqlgen/graphql"
	"github.com/brandon-julio-t/graph-gongular-backend/graph/model"
	"github.com/brandon-julio-t/graph-gongular-backend/middlewares"
	"github.com/google/uuid"
)

func (r *mutationResolver) Upload(ctx context.Context, files []*graphql.Upload) (bool, error) {
	user := middlewares.UseAuth(ctx)
	if user == nil {
		return false, errors.New("not authenticated")
	}

	for _, file := range files {
		filenameSplit := strings.Split(file.Filename, ".")
		extension := filenameSplit[len(filenameSplit)-1]
		baseDir := "storage"

		id := uuid.Must(uuid.NewRandom()).String()
		filename := fmt.Sprintf("%v.%v", id, extension)

		path := filepath.Join(baseDir, filename)

		if _, err := os.Stat(baseDir); os.IsNotExist(err) {
			log.Printf("path %v doesn't exists. creating it...", baseDir)
			if err := os.Mkdir(baseDir, os.ModePerm); err != nil {
				return false, err
			}
		}

		data, err := ioutil.ReadAll(file.File)
		if len(data) < 1 {
			return false, fmt.Errorf("file is empty %v\n", file)
		} else if err != nil {
			return false, err
		}

		if err := ioutil.WriteFile(path, data, os.ModePerm); err != nil {
			return false, err
		}

		saved, err := r.UserService.SaveFile(&model.FileUpload{
			ID:          id,
			Path:        path,
			Filename:    file.Filename,
			Size:        int(file.Size),
			ContentType: file.ContentType,
			UserID:      user.ID,
		})

		if err != nil {
			return false, fmt.Errorf("cannot save file %v\n", saved)
		}
	}

	return true, nil
}
