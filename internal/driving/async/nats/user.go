package nats

import (
	"gitlab.com/ricardo134/auth-service/internal/core/ports/user"
)

type UserHandler interface {
	Requested(userID uint)
}

type userHandler struct {
	userService user.Service
}

func NewNatsUserHandler(userSvc user.Service) UserHandler {
	return userHandler{userSvc}
}

func (nh userHandler) Requested(userID uint) {
	_, _ = nh.userService.Get(userID)
}
