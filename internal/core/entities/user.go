package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"uniqueIndex, notNull"`
	Email    string `json:"email" gorm:"uniqueIndex, notNull"`
	Password string `json:"password" gorm:"notNull"`
}

type CreateUserRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// LoginRequest It is similar to CreateUserRequest for now, but the other struct is going to change in the future
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
