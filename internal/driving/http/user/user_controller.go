package user

import (
	"github.com/gin-gonic/gin"
	ricardoErr "gitlab.com/ricardo-public/errors/pkg/errors"
	tokens "gitlab.com/ricardo-public/jwt-tools/v2/pkg/token"
	"gitlab.com/ricardo134/auth-service/internal/core/app/user"
	"gitlab.com/ricardo134/auth-service/internal/core/entities"
	"net/http"
	"strconv"
)

type Controller interface {
	Get(gtx *gin.Context)
	Me(gtx *gin.Context)
	Update(gtx *gin.Context)
	Delete(gtx *gin.Context)
}

func NewController(service user.Service) Controller {
	return controller{uSvc: service}
}

type controller struct {
	uSvc user.Service
}

// Get
// @Summary Get user
// @Description Get a user
// @Success 200 {object} entities.User
// @Failure 400 {object} ricardoErr.RicardoError
// @Failure 404 {object} ricardoErr.RicardoError
// @Router /users/{user_id} [get]
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

// Me
// @Summary Get connected user
// @Description Get connected user
// @Success 200 {object} entities.User
// @Failure 404 {object} ricardoErr.RicardoError
// @Router /users/me [get]
func (c controller) Me(gtx *gin.Context) {
	userID, _ := gtx.Get(tokens.UserIDKey)
	user, err := c.uSvc.Get(userID.(uint))
	if err != nil {
		_ = ricardoErr.GinErrorHandler(gtx, err)
		return
	}

	gtx.JSON(http.StatusOK, *user)
}

// Update
// @Summary Update a user
// @Description Update a user
// @Param user_id path int true "User id"
// @Param user body entities.UpdateUserRequest true "Updated user info"
// @Success 200 {object} entities.User
// @Failure 400 {object} ricardoErr.RicardoError
// @Failure 404 {object} ricardoErr.RicardoError
// @Router /users/{user_id} [patch]
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
		ID:          uint(userID),
		Username:    updUserReq.Username,
		Email:       updUserReq.Email,
		IsSetupDone: updUserReq.IsSetupDone,
	}
	updUser, err := c.uSvc.Update(u)
	if err != nil {
		_ = ricardoErr.GinErrorHandler(gtx, err)
		return
	}

	gtx.JSON(http.StatusOK, *updUser)
}

// Delete
// @Summary Delete a user
// @Description Delete a user
// @Param user_id path int true "User id"
// @Success 200 {object} entities.User
// @Failure 400 {object} ricardoErr.RicardoError
// @Failure 404 {object} ricardoErr.RicardoError
// @Router /users/{user_id} [delete]
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
