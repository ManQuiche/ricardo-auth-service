package boot

import (
	"fmt"
	"log"
	"net/http"
	"ricardo/auth-service/internal/driving/http/auth"

	"github.com/gin-gonic/gin"
)

var (
	router *gin.Engine
)

// @title auth-service
// @version 1.0
// @description Ricardo's auth service.
//
// @accept json
// @produce json
//
// @contact.name   Ricardo teams
// @contact.email  support@ricardo.net
//
// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
//
// @BasePath  /auth

func initRoutes() {
	// Ready route
	router.GET("/", func(context *gin.Context) {
		context.Status(http.StatusOK)
	})

	authrController := auth.NewBasicController(authenticateService)
	firebaseController := auth.NewFirebaseController(externalTokenService)

	authRouter := router.Group("/auth")
	authRouter.POST("/login", authrController.Login)
	authRouter.POST("/register", authrController.Create)

	// TODO: add a firebase controller
	firebaseRouter := authRouter.Group("/firebase")
	firebaseRouter.POST("/login", firebaseController.Login)

	// JWT Middleware definition
	accessMiddleware := auth.NewJwtAuthMiddleware(authorizationService, false)
	refreshMiddleware := auth.NewJwtAuthMiddleware(authorizationService, true)

	// @Summary Serve a new token pair
	// @Description Serve a new refresh token if the given one is good, and a new access token too
	// @Success 200
	// @Failure 401
	// @Router /auth/access [post]
	authRouter.GET("/access", accessMiddleware.Authorize, func(context *gin.Context) {
		context.Status(http.StatusOK)
	})

	authRouter.GET("/refresh", refreshMiddleware.Authorize, authrController.Refresh)
}

func ServeHTTP() {
	router = gin.Default()

	initRoutes()

	appURL := fmt.Sprintf("%s:%s", url, port)
	log.Printf("Launching server on %s...\n", appURL)

	log.Fatalln(router.Run(appURL))

	// TODO: go func and etc
	//log.Println("HTTP server stopped, exiting...")
}
