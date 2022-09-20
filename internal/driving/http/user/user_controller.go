package user

import (
	"github.com/gin-gonic/gin"
	ricardoErr "gitlab.com/ricardo-public/errors/pkg/errors"
	"gitlab.com/ricardo134/auth-service/internal/core/app/user"
	"gitlab.com/ricardo134/auth-service/internal/core/entities"
	"gorm.io/gorm"
	"net/http"
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
	var getUserReq entities.GetUserRequest
	err := gtx.ShouldBindJSON(&getUserReq)
	if err != nil {
		_ = ricardoErr.GinErrorHandler(gtx, ricardoErr.New(ricardoErr.ErrBadRequest, err.Error()))
		return
	}

	gUser, err := c.uSvc.Get(getUserReq.ID)
	if err != nil {
		_ = ricardoErr.GinErrorHandler(gtx, err)
		return
	}

	gtx.JSON(http.StatusOK, *gUser)
}

func (c controller) Update(gtx *gin.Context) {
	var updUserReq entities.UpdateUserRequest
	err := gtx.ShouldBindJSON(&updUserReq)
	if err != nil {
		_ = ricardoErr.GinErrorHandler(gtx, ricardoErr.New(ricardoErr.ErrBadRequest, err.Error()))
		return
	}

	u := entities.User{
		Model: gorm.Model{
			ID: updUserReq.ID,
		},
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
	var delUserReq entities.DeleteUserRequest
	err := gtx.ShouldBindJSON(&delUserReq)
	if err != nil {
		_ = ricardoErr.GinErrorHandler(gtx, ricardoErr.New(ricardoErr.ErrBadRequest, err.Error()))
		return
	}

	dUser, err := c.uSvc.Delete(delUserReq.ID)
	if err != nil {
		_ = ricardoErr.GinErrorHandler(gtx, err)
		return
	}

	gtx.JSON(http.StatusOK, dUser)
}
