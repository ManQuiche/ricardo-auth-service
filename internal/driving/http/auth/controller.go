package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ricardo/auth-service/internal/core/app/auth"
	"ricardo/auth-service/internal/core/entities"
	errors2 "ricardo/auth-service/pkg/errors"
)

type Controller interface {
	Create(gtx *gin.Context)
	Login(gtx *gin.Context)
	Refresh(gtx *gin.Context)
}

func NewController(service auth.AuthenticateService) Controller {
	return controller{authr: service}
}

type controller struct {
	authr auth.AuthenticateService
}

func (c controller) Refresh(gtx *gin.Context) {
	// invalidate old ?

	// get new access token
}

func (c controller) Create(gtx *gin.Context) {
	var createUserRequest entities.CreateUserRequest
	err := gtx.ShouldBindJSON(&createUserRequest)
	if err != nil {
		_ = errors2.GinErrorHandler(gtx, err, http.StatusBadRequest)
		return
	}

	user := entities.User{
		Username: createUserRequest.Username,
		Email:    createUserRequest.Email,
		Password: createUserRequest.Password,
	}
	err = c.authr.Save(gtx.Request.Context(), user)
	if err != nil {
		_ = errors2.GinErrorHandler(gtx, err, http.StatusInternalServerError)
		return
	}
}

func (c controller) Login(gtx *gin.Context) {
	var loginRequest entities.LoginRequest
	err := gtx.ShouldBindJSON(&loginRequest)
	if err != nil {
		_ = errors2.GinErrorHandler(gtx, err, http.StatusBadRequest)
		return
	}

	tokens, err := c.authr.Login(gtx.Request.Context(), loginRequest)
	if err != nil {
		_ = errors2.GinErrorHandler(gtx, err, http.StatusNotFound)
		return
	}

	gtx.JSON(http.StatusOK, tokens)
}
