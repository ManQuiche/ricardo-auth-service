package main

import (
	authboot "auth-service/boot"
	"auth-service/internal/auth/firebase"
	authhttp "auth-service/internal/http"
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

	authboot.LoadEnv()

	firebase.InitFirebaseSDK()

	router := httprouter.New()
	authhttp.InitRoute(router)

	appURL := fmt.Sprintf("%s:%s", os.Getenv("URL"), os.Getenv("PORT"))
	log.Printf("Launching server on %s...\n", appURL)
	log.Fatal(http.ListenAndServe(appURL, router))
}
