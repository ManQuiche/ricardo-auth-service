package auth

import (
	"auth-service/internal/core/app/auth"
	autherrors "auth-service/pkg/errors"
	authhttp "auth-service/pkg/http"
	"errors"
	"github.com/gin-gonic/gin"
	"log"
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
	token := gtx.GetHeader(authhttp.AuthorizationHeader)[len(authhttp.BearerType):]
	log.Println(token)

	if token == "" {
		gtx.AbortWithStatus(http.StatusUnauthorized)
	}

	if j.refresh {
		authorized, err = j.authz.RefreshAuthorize(gtx.Request.Context(), token)
	} else {
		authorized, err = j.authz.AccessAuthorize(gtx.Request.Context(), token)
	}

	if authorized {
		gtx.Next()
	} else if !authorized {
		_ = autherrors.GinErrorHandler(gtx, errors.New("access token is invalid"), http.StatusUnauthorized)
	} else if err != nil {
		_ = autherrors.GinErrorHandler(gtx, err, http.StatusInternalServerError)
	}
}
