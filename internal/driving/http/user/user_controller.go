package user

import (
	"github.com/gin-gonic/gin"
	ricardoErr "gitlab.com/ricardo-public/errors/pkg/errors"
	"gitlab.com/ricardo134/auth-service/internal/core/app/user"
	"gitlab.com/ricardo134/auth-service/internal/core/entities"
	"net/http"
	"strconv"
)

type Controller interface {
	Get(gtx *gin.Context)
	Update(gtx *gin.Context)
	Delete(gtx *gin.Context)
}

func NewController(service user.Service) Controller {
	return controller{uSvc: service}
}

type controller struct {
	uSvc user.Service
}

func (c controller) Get(gtx *gin.Context) {
	userID, err := strconv.Atoi(gtx.Param("user_id"))
	if err != nil {
		_ = ricardoErr.GinErrorHandler(gtx, ricardoErr.New(ricardoErr.ErrBadRequest, err.Error()))
		return
	}

	gUser, err := c.uSvc.Get(uint(userID))
	if err != nil {
		_ = ricardoErr.GinErrorHandler(gtx, err)
		return
	}

	gtx.JSON(http.StatusOK, *gUser)
}

func (c controller) Update(gtx *gin.Context) {
	userID, err := strconv.Atoi(gtx.Param("user_id"))
	if err != nil {
		_ = ricardoErr.GinErrorHandler(gtx, ricardoErr.New(ricardoErr.ErrBadRequest, err.Error()))
		return
	}

	var updUserReq entities.UpdateUserRequest
	err = gtx.ShouldBindJSON(&updUserReq)
	if err != nil {
		_ = ricardoErr.GinErrorHandler(gtx, ricardoErr.New(ricardoErr.ErrBadRequest, err.Error()))
		return
	}

	u := entities.User{
		ID:       uint(userID),
		Username: updUserReq.Username,
		Email:    updUserReq.Email,
	}
	updUser, err := c.uSvc.Update(u)
	if err != nil {
		_ = ricardoErr.GinErrorHandler(gtx, err)
		return
	}

	gtx.JSON(http.StatusOK, *updUser)
}

func (c controller) Delete(gtx *gin.Context) {
	userID, err := strconv.Atoi(gtx.Param("user_id"))
	if err != nil {
		_ = ricardoErr.GinErrorHandler(gtx, ricardoErr.New(ricardoErr.ErrBadRequest, err.Error()))
		return
	}

	dUser, err := c.uSvc.Delete(uint(userID))
	if err != nil {
		_ = ricardoErr.GinErrorHandler(gtx, err)
		return
	}

	gtx.JSON(http.StatusOK, dUser)
}
