package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/99designs/gqlgen/graphql"
	"github.com/brandon-julio-t/graph-gongular-backend/facades"
	"github.com/brandon-julio-t/graph-gongular-backend/graph/model"
	"github.com/brandon-julio-t/graph-gongular-backend/middlewares"
)

func (r *mutationResolver) UpdateFile(ctx context.Context, input *model.UpdateFile) (*model.FileUpload, error) {
	if user := middlewares.UseAuth(ctx); user == nil {
		return nil, facades.NotAuthenticatedError
	}
	return r.FileUploadService.UpdateFile(input)
}

func (r *mutationResolver) DeleteFile(ctx context.Context, id string) (*model.FileUpload, error) {
	if user := middlewares.UseAuth(ctx); user == nil {
		return nil, facades.NotAuthenticatedError
	}
	return r.FileUploadService.DeleteFile(id)
}

func (r *mutationResolver) Upload(ctx context.Context, files []*graphql.Upload) (bool, error) {
	user := middlewares.UseAuth(ctx)
	if user == nil {
		return false, facades.NotAuthenticatedError
	}

	for _, file := range files {
		if saved, err := r.FileUploadService.SaveFile(file, user); err != nil {
			return false, fmt.Errorf("cannot save file %v\n", saved)
		}
	}

	return true, nil
}

func (r *queryResolver) Files(ctx context.Context) ([]*model.FileUpload, error) {
	user := middlewares.UseAuth(ctx)
	if user == nil {
		return nil, facades.NotAuthenticatedError
	}
	return r.FileUploadService.GetFilesByUser(user)
}

func (r *queryResolver) Download(ctx context.Context, id string) (string, error) {
	user := middlewares.UseAuth(ctx)
	if user == nil {
		return "", facades.NotAuthenticatedError
	}

	file, err := r.FileUploadService.GetFileById(id)
	if err != nil {
		return "", err
	}

	if file.UserID != user.ID {
		return "", errors.New("unauthorized download")
	}

	filename := fmt.Sprintf("%v.%v", file.ID, file.Extension)
	path := filepath.Join(facades.BaseDir, filename)
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(data), nil
}
