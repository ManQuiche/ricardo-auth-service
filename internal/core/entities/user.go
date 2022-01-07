package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username,omitempty" gorm:"index"`
	Password string `json:"password,omitempty"`
}

type CreateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
