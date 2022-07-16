package auth

import (
	"errors"
	ricardoErr "gitlab.com/ricardo-public/errors/pkg/errors"
	"net/http"
	"ricardo/auth-service/internal/core/app/auth"
	authhttp "ricardo/auth-service/pkg/http"

	"github.com/gin-gonic/gin"
)

type FirebaseController interface {
	Login(gtx *gin.Context)
}

func NewFirebaseController(service auth.ExternalTokenService) FirebaseController {
	return firebaseController{extService: service}
}

type firebaseController struct {
	extService auth.ExternalTokenService
}

// Login
// @Summary Login
// @Description Login from a firebase token
// @Success 200 {object} entities.SignedTokenPair
// @Failure 400 {object} ricardoErr.RicardoError
// @Failure 404 {object} ricardoErr.RicardoError
// @Router /auth/firebase/login [post]
func (f firebaseController) Login(gtx *gin.Context) {
	var err error
	token := gtx.GetHeader(authhttp.AuthorizationHeader)
	if token == "" {
		ricardoErr.GinErrorHandlerWithCode(gtx, errors.New("you need to pass an access token"), http.StatusUnauthorized)
		return
	}

	if len(token) <= len(authhttp.BearerType) {
		ricardoErr.GinErrorHandlerWithCode(gtx, errors.New("access token format is invalid"), http.StatusUnauthorized)
		return
	}
	token = token[len(authhttp.BearerType):]

	tokens, err := f.extService.Verify(gtx, token)
	if err != nil {
		ricardoErr.GinErrorHandler(gtx, err)
		return
	}
	gtx.JSON(http.StatusOK, tokens)
}
