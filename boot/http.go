package boot

import (
	"auth-service/internal/driving/http/auth"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

var (
	router *gin.Engine
)

func initRoutes() {
	authrController := auth.NewController(authenticateService)

	router.POST("/login", authrController.Login)
	userGroup := router.Group("/users")
	userGroup.POST("", authrController.Create)

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
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
