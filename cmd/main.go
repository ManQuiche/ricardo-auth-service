package main

import (
	"auth-service/boot"
	authhttp "auth-service/internal/http"
	"auth-service/internal/storage/firebase"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

func main() {
	defer func() {
		log.Println("Exiting...")
	}()

	boot.LoadEnv()
	boot.LoadDb()

	firebase.InitFirebaseSDK()

	router := httprouter.New()
	authhttp.InitRoute(router)

	appURL := fmt.Sprintf("%s:%s", os.Getenv("URL"), os.Getenv("PORT"))
	log.Printf("Launching server on %s...\n", appURL)
	log.Fatal(http.ListenAndServe(appURL, router))
}
