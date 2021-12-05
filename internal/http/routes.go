package http

import (
	"auth-service/internal/auth"
	"github.com/julienschmidt/httprouter"
)

func InitRoute(router *httprouter.Router) *httprouter.Router {
	router.POST("/create", auth.CreateUser)
	router.POST("/login", auth.LoginUser)

	return router
}
