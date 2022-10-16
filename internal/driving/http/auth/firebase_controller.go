package auth

import (
	"errors"
	errorsext "gitlab.com/ricardo-public/errors/pkg/errors"
	_ "gitlab.com/ricardo-public/jwt-tools/v2/pkg/token"
	"gitlab.com/ricardo-public/tracing/pkg/tracing"
	"gitlab.com/ricardo134/auth-service/internal/core/app/auth"
	authhttp "gitlab.com/ricardo134/auth-service/pkg/http"
	"go.opentelemetry.io/otel/codes"
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"
	"go.opentelemetry.io/otel/trace"
	"net/http"
	"strconv"

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
// @Success 200 {object} token.SignedTokens
// @Failure 400 {object} errorsext.RicardoError
// @Failure 404 {object} errorsext.RicardoError
// @Router /auth/firebase/login [post]
func (f firebaseController) Login(gtx *gin.Context) {
	span := gtx.Value(tracing.HttpSpanKey).(trace.Span)
	defer span.End()

	var err error
	token := gtx.GetHeader(authhttp.AuthorizationHeader)
	if token == "" {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		span.SetAttributes(semconv.HTTPStatusCodeKey.String(strconv.Itoa(http.StatusUnauthorized)))
		errorsext.GinErrorHandlerWithCode(gtx, errors.New("you need to pass an access token"), http.StatusUnauthorized)
		return
	}

	if len(token) <= len(authhttp.BearerType) {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		span.SetAttributes(semconv.HTTPStatusCodeKey.String(strconv.Itoa(http.StatusUnauthorized)))
		errorsext.GinErrorHandlerWithCode(gtx, errors.New("access token format is invalid"), http.StatusUnauthorized)
		return
	}
	token = token[len(authhttp.BearerType):]

	tokens, err := f.extService.Verify(gtx, token)
	if err != nil {
		span.SetAttributes(semconv.HTTPStatusCodeKey.String(strconv.Itoa(http.StatusUnauthorized)))
		errorsext.GinErrorHandler(gtx, err)
		return
	}

	span.SetAttributes(semconv.HTTPStatusCodeKey.String(strconv.Itoa(http.StatusOK)))
	gtx.JSON(http.StatusOK, tokens)
}
