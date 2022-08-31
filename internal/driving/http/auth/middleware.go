package auth

import (
	"errors"
	"github.com/gin-gonic/gin"
	ricardoErr "gitlab.com/ricardo-public/errors/pkg/errors"
	"gitlab.com/ricardo134/auth-service/internal/core/app/auth"
	authhttp "gitlab.com/ricardo134/auth-service/pkg/http"
	"net/http"
)

type AuthMiddleware interface {
	Authorize(c *gin.Context)
}

type jwtAuthMiddleware struct {
	authz   auth.AuthorizeService
	refresh bool
}

func NewJwtAuthMiddleware(authz auth.AuthorizeService, refresh bool) AuthMiddleware {
	return jwtAuthMiddleware{
		authz:   authz,
		refresh: refresh,
	}
}

func (j jwtAuthMiddleware) Authorize(gtx *gin.Context) {
	authorized := false
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

	if j.refresh {
		authorized, err = j.authz.RefreshAuthorize(gtx.Request.Context(), token)
	} else {
		authorized, err = j.authz.AccessAuthorize(gtx.Request.Context(), token)
	}

	if authorized {
		gtx.Next()
	} else if !authorized {
		_ = ricardoErr.GinErrorHandlerWithCode(gtx, errors.New("access token is invalid"), http.StatusUnauthorized)
	} else if err != nil {
		_ = ricardoErr.GinErrorHandlerWithCode(gtx, err, http.StatusInternalServerError)
	}
}
