package boot

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

var (
	router *gin.Engine
)

func initRoutes() {
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

	go func() {
		log.Fatalln(router.Run(appURL))
	}()

	log.Println("HTTP server stopped, exiting...")
}
