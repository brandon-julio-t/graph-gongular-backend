package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/brandon-julio-t/graph-gongular-backend/graph/model"
	"github.com/brandon-julio-t/graph-gongular-backend/middlewares"
)

func (r *mutationResolver) UpdateFile(ctx context.Context, input *model.UpdateFile) (*model.FileUpload, error) {
	if user := middlewares.UseAuth(ctx); user == nil {
		return nil, errors.New("not authenticated")
	}
	return r.UserService.UpdateFile(input)
}

func (r *mutationResolver) DeleteFile(ctx context.Context, id string) (*model.FileUpload, error) {
	if user := middlewares.UseAuth(ctx); user == nil {
		return nil, errors.New("not authenticated")
	}

	file, err := r.UserService.DeleteFile(id)
	if err != nil {
		return nil, err
	}

	filenameSplit := strings.Split(file.Filename, ".")
	extension := filenameSplit[len(filenameSplit)-1]
	filename := fmt.Sprintf("%v.%v", file.ID, extension)
	path := filepath.Join("storage", filename)
	if err := os.Remove(path); err != nil {
		return nil, err
	}

	return file, err
}
