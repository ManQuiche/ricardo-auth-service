package boot

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"ricardo/auth-service/internal/driving/http/auth"
)

var (
	router *gin.Engine
)

func initRoutes() {
	// Ready route
	router.GET("/", func(context *gin.Context) {
		context.Status(http.StatusOK)
	})

	authrController := auth.NewController(authenticateService)

	authRouter := router.Group("/auth")
	authRouter.POST("/login", authrController.Login)
	authRouter.POST("/register", authrController.Create)

	// JWT Middleware definition
	accessMiddleware := auth.NewJwtAuthMiddleware(authorizationService, false)
	refreshMiddleware := auth.NewJwtAuthMiddleware(authorizationService, true)

	// Access Token check route
	authRouter.GET("/access", accessMiddleware.Authorize, func(context *gin.Context) {
		context.Status(http.StatusOK)
	})

	authRouter.GET("/refresh", refreshMiddleware.Authorize, authrController.Refresh)

	//userGroup := router.Group("/users", accessMiddleware.Authorize)
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
