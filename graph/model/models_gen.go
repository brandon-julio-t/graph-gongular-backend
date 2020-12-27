// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"
)

type FileUpload struct {
	ID          string `json:"id"`
	Path        string `json:"path"`
	Filename    string `json:"filename"`
	Size        int    `json:"size"`
	ContentType string `json:"contentType"`
	UserID      string `json:"userId"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Register struct {
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	DateOfBirth time.Time `json:"dateOfBirth"`
	Gender      string    `json:"gender"`
	Address     string    `json:"address"`
}

type UpdateFile struct {
	ID       string `json:"id"`
	Filename string `json:"filename"`
}

type UpdateUser struct {
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	DateOfBirth time.Time `json:"dateOfBirth"`
	Gender      string    `json:"gender"`
	Address     string    `json:"address"`
}

type User struct {
	ID          string        `json:"id"`
	Name        string        `json:"name"`
	Email       string        `json:"email"`
	Password    string        `json:"password"`
	DateOfBirth time.Time     `json:"dateOfBirth"`
	Gender      string        `json:"gender"`
	Address     string        `json:"address"`
	UserRoleID  string        `json:"userRoleId"`
	UserRole    *UserRole     `json:"userRole"`
	Files       []*FileUpload `json:"files"`
}

type UserRole struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
