package model

type FileUpload struct {
	ID          string `json:"id"`
	File        []byte `json:"file"`
	Filename    string `json:"filename"`
	Extension   string `json:"extension"`
	Size        int64  `json:"size"`
	ContentType string `json:"contentType"`
	UserID      string `json:"userId"`
	User        *User  `json:"user"`
}
