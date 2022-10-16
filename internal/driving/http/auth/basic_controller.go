package auth

import (
	"gitlab.com/ricardo-public/tracing/pkg/tracing"
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"
	"net/http"
	"strconv"

	"gitlab.com/ricardo134/auth-service/internal/core/app/auth"
	"gitlab.com/ricardo134/auth-service/internal/core/entities"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"

	"github.com/gin-gonic/gin"
	errorsext "gitlab.com/ricardo-public/errors/pkg/errors"
	token "gitlab.com/ricardo-public/jwt-tools/v2/pkg/token"
)

type BasicController interface {
	Create(gtx *gin.Context)
	Login(gtx *gin.Context)
	Refresh(gtx *gin.Context)
}

func NewBasicController(service auth.AuthenticateService) BasicController {
	return basicController{authr: service}
}

type basicController struct {
	authr auth.AuthenticateService
}

// Refresh
// @Summary Serve a new token pair
// @Description Serve a new refresh token if the given one is good, and a new access token too
// @Success 200 {object} token.SignedTokens
// @Failure 401 {object} errorsext.RicardoError
// @Router /auth/refresh [post]
func (c basicController) Refresh(gtx *gin.Context) {
	// TODO: invalidate old token pair
	span := gtx.Request.Context().Value(tracing.HttpSpanKey).(trace.Span)
	var err error
	defer span.End()

	// there will be no error since the token has already been checked in the middleware
	token, _ := token.ExtractTokenFromHeader(gtx.GetHeader(token.AuthorizationHeader))

	tokenPair, err := c.authr.Refresh(gtx.Request.Context(), token)
	if err != nil {
		span.SetAttributes(semconv.HTTPStatusCodeKey.String(strconv.Itoa(http.StatusInternalServerError)))
		_ = errorsext.GinErrorHandler(gtx, errorsext.New("TODO: add internal server error type to error lib", err.Error()))
		return
	}

	gtx.JSON(http.StatusOK, tokenPair)
}

// Create
// @Summary Create a new user
// @Description Create a new user
// @Success 200
// @Failure 400 {object} errorsext.RicardoError
// @Failure 403 {object} errorsext.RicardoError
// @Router /auth/register [post]
func (c basicController) Create(gtx *gin.Context) {
	span := gtx.Request.Context().Value(tracing.HttpSpanKey).(trace.Span)
	var err error
	defer span.End()

	var createUserRequest entities.CreateUserRequest
	err = gtx.ShouldBindJSON(&createUserRequest)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		span.SetAttributes(semconv.HTTPStatusCodeKey.String(strconv.Itoa(http.StatusBadRequest)))
		_ = errorsext.GinErrorHandler(gtx, errorsext.New(errorsext.ErrBadRequest, err.Error()))
		return
	}

	user := entities.User{
		Username: createUserRequest.Username,
		Email:    createUserRequest.Email,
		Password: createUserRequest.Password,
	}
	err = c.authr.Save(gtx.Request.Context(), user)
	if err != nil {
		span.SetAttributes(semconv.HTTPStatusCodeKey.String(strconv.Itoa(http.StatusInternalServerError)))
		_ = errorsext.GinErrorHandler(gtx, err)
		return
	}
}

// Login
// @Summary Login
// @Description Login from an email and a password
// @Success 200 {object} token.SignedTokens
// @Failure 400 {object} errorsext.RicardoError
// @Failure 404 {object} errorsext.RicardoError
// @Router /auth/login [post]
func (c basicController) Login(gtx *gin.Context) {
	span := gtx.Request.Context().Value(tracing.HttpSpanKey).(trace.Span)
	var err error
	defer span.End()

	var loginRequest entities.LoginRequest
	err = gtx.ShouldBindJSON(&loginRequest)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		span.SetAttributes(semconv.HTTPStatusCodeKey.String(strconv.Itoa(http.StatusBadRequest)))
		_ = errorsext.GinErrorHandler(gtx, errorsext.New(errorsext.ErrBadRequest, err.Error()))
		return
	}

	tokens, err := c.authr.Login(gtx.Request.Context(), loginRequest)
	if err != nil {
		span.SetAttributes(semconv.HTTPStatusCodeKey.String(strconv.Itoa(http.StatusUnauthorized)))
		_ = errorsext.GinErrorHandler(gtx, err)
		return
	}

	span.SetAttributes(semconv.HTTPStatusCodeKey.String(strconv.Itoa(http.StatusOK)))
	gtx.JSON(http.StatusOK, tokens)
}
