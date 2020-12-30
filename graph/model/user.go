package model

import "time"

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
	Friends     []*User       `json:"friends" gorm:"many2many:friends;"`
	FileUploads []*FileUpload `json:"fileUploads"`
}
