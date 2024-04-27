package boot

import (
	"fmt"
	tokens "gitlab.com/ricardo-public/jwt-tools/v2/pkg/token"
	gintracing "gitlab.com/ricardo-public/tracing/pkg/gin"
	"gitlab.com/ricardo134/auth-service/internal/driving/http/auth"
	"gitlab.com/ricardo134/auth-service/internal/driving/http/user"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	router *gin.Engine
)

const (
	SpanKey = "span"
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

func initRoutes() {
	router.Use(gintracing.TraceRequest)

	// Ready route
	router.GET("/", func(context *gin.Context) {
		context.Status(http.StatusOK)
	})

	accessMiddleware := tokens.NewJwtAuthMiddleware([]byte(accessSecret))
	refreshMiddleware := tokens.NewJwtAuthMiddleware([]byte(refreshSecret))

	authrController := auth.NewBasicController(authenticateService)
	firebaseController := auth.NewFirebaseController(externalTokenService)
	userController := user.NewController(userService)

	authRouter := router.Group("/auth")
	authRouter.POST("/login", authrController.Login)
	authRouter.POST("/register", authrController.Create)
	// @Summary Serve a new token pair
	// @Description Serve a new refresh token if the given one is good, and a new access token too
	// @Success 200
	// @Failure 401
	// @Router /auth/access [post]
	authRouter.GET("/access", accessMiddleware.Authorize, func(context *gin.Context) {
		context.Status(http.StatusOK)
	})
	authRouter.GET("/refresh", refreshMiddleware.Authorize, authrController.Refresh)

	firebaseRouter := authRouter.Group("/firebase")
	firebaseRouter.POST("/login", firebaseController.Login)

	usrRouter := router.Group("/users", accessMiddleware.Authorize)
	usrRouter.GET("/:user_id", userController.Get)
	usrRouter.GET("/me", userController.Me)
	usrRouter.PUT("/:user_id", userController.Update)
	usrRouter.DELETE("/:user_id", userController.Delete)
}

func ServeHTTP() {
	router = gin.Default()
	_ = router.SetTrustedProxies(nil)

	initRoutes()

	appURL := fmt.Sprintf("%s:%s", url, port)
	log.Printf("Launching server on %s...\n", appURL)

	log.Fatalln(router.Run(appURL))

	// TODO: go func and etc
	//log.Println("HTTP server stopped, exiting...")
}
