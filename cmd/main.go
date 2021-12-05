package main

import (
	"auth-service/internal/auth/firebase"
	authhttp "auth-service/internal/http"
	"auth-service/pkg/errors"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"os"
)

func checkEnv(name string) {
	if os.Getenv(name) == "" {
		errors.MissingEnvVarF(name)
	}
}

func init() {
	checkEnv("URL")
	checkEnv("PORT")
}

func main() {
	defer func() {
		log.Println("Exiting...")
	}()

	firebase.InitFirebaseSDK()

	router := httprouter.New()
	authhttp.InitRoute(router)

	appURL := fmt.Sprintf("%s:%s", os.Getenv("URL"), os.Getenv("PORT"))
	log.Printf("Launching server on %s...\n", appURL)
	log.Fatal(http.ListenAndServe(appURL, router))
}
