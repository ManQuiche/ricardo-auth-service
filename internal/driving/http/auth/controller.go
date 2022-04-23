package auth

import (
	"github.com/gin-gonic/gin"
	ricardoErr "gitlab.com/ricardo-public/errors/pkg/errors"
	"net/http"
	"ricardo/auth-service/internal/core/app/auth"
	"ricardo/auth-service/internal/core/entities"
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

// Refresh
// @Summary Serve a new token pair
// @Description Serve a new refresh token if the given one is good, and a new access token too
// @Success 200 {object} entities.SignedTokenPair
// @Failure 400 {object} ricardoErr.RicardoError
// @Router /auth/refresh [post]
func (c controller) Refresh(gtx *gin.Context) {
	// invalidate old ?

	// get new access token
}

func (c controller) Create(gtx *gin.Context) {
	var createUserRequest entities.CreateUserRequest
	err := gtx.ShouldBindJSON(&createUserRequest)
	if err != nil {
		_ = ricardoErr.GinErrorHandler(gtx, ricardoErr.New(ricardoErr.ErrBadRequest, err.Error()))
		return
	}

	user := entities.User{
		Username: createUserRequest.Username,
		Email:    createUserRequest.Email,
		Password: createUserRequest.Password,
	}
	err = c.authr.Save(gtx.Request.Context(), user)
	if err != nil {
		_ = ricardoErr.GinErrorHandler(gtx, ricardoErr.New("TODO: add internal server error type to error lib", err.Error()))
		return
	}
}

// Login
// @Summary Login
// @Description Login from an email and a password
// @Success 200 {object} entities.SignedTokenPair
// @Failure 400 {object} ricardoErr.RicardoError
// @Failure 404 {object} ricardoErr.RicardoError
// @Router /auth/login [post]
func (c controller) Login(gtx *gin.Context) {
	var loginRequest entities.LoginRequest
	err := gtx.ShouldBindJSON(&loginRequest)
	if err != nil {
		_ = ricardoErr.GinErrorHandler(gtx, ricardoErr.New(ricardoErr.ErrBadRequest, err.Error()))
		return
	}

	tokens, err := c.authr.Login(gtx.Request.Context(), loginRequest)
	if err != nil {
		_ = ricardoErr.GinErrorHandler(gtx, ricardoErr.New("TODO: add internal server error type to error lib", err.Error()))
		return
	}

	gtx.JSON(http.StatusOK, tokens)
}
