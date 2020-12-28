package model

import "github.com/99designs/gqlgen/graphql"

type FileUpload struct {
	*graphql.Upload
	ID          string `json:"id"`
	Path        string `json:"path"`
	Extension   string `json:"extension"`
	UserID      string `json:"userId"`
}
