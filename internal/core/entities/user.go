package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"uniqueIndex, notNull"`
	Password string `json:"password" gorm:"notNull"`
}

type CreateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// LoginRequest It is similar to CreateUserRequest for now, but the other struct is going to change in the future
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
