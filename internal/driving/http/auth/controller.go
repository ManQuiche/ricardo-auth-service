package auth

import (
	"auth-service/internal/core/app/auth"
	"auth-service/internal/core/entities"
	errors2 "auth-service/pkg/errors"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller interface {
	Create(gtx *gin.Context)
	Login(gtx *gin.Context)
}

func NewController(service auth.AuthenticateService) Controller {
	return controller{authr: service}
}

type controller struct {
	authr auth.AuthenticateService
}

func (c controller) Create(gtx *gin.Context) {
	var createUserRequest entities.CreateUserRequest
	err := gtx.BindJSON(&createUserRequest)
	if err != nil {
		_ = errors2.GinErrorHandler(gtx, err, http.StatusBadRequest)
		return
	}

	if createUserRequest.Username == "" || createUserRequest.Password == "" {
		_ = errors2.GinErrorHandler(gtx, errors.New("username and password cannot be null"), http.StatusBadRequest)
		return
	}

	user := entities.User{
		Username: createUserRequest.Username,
		Password: createUserRequest.Password,
	}
	err = c.authr.Save(gtx.Request.Context(), user)
	if err != nil {
		_ = errors2.GinErrorHandler(gtx, err, http.StatusInternalServerError)
		return
	}
}

func (controller) Login(gtx *gin.Context) {
	//TODO implement me
	panic("implement me")
}
