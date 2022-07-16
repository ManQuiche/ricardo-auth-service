package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"uniqueIndex"`
	Email    string `json:"email" gorm:"uniqueIndex, notNull"`
	Password string `json:"password" gorm:"notNull"`
	// ExternalSource Must be lowercase
	ExternalSource string `json:"external_source"`
}

type CreateUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// LoginRequest It is similar to CreateUserRequest for now, but the other struct is going to change in the future
type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
