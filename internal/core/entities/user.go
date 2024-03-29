package entities

import (
	tokens "gitlab.com/ricardo-public/jwt-tools/v2/pkg/token"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID          uint        `json:"id" gorm:"primarykey"`
	Username    string      `json:"username"`
	Email       string      `json:"email" gorm:"uniqueIndex, notNull"`
	Password    string      `json:"-" gorm:"notNull"`
	Role        tokens.Role `json:"role" gorm:"notNull,type:string"`
	IsSetupDone bool        `json:"is_setup_done" gorm:"notNull,default:false"`
	// ExternalSource Must be lowercase
	ExternalSource string `json:"external_source"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

type ShortUser struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
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

type UpdateUserRequest struct {
	Username    string `json:"username" binding:"required"`
	Email       string `json:"email" binding:"required"`
	IsSetupDone bool   `json:"is_setup_done" binding:"required"`
}
