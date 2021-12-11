package main

import (
	"auth-service/boot"
	"auth-service/internal/storage/firebase"
)

func main() {
	boot.LoadEnv()
	boot.LoadDb()

	firebase.InitFirebaseSDK()

	boot.ServeHTTP()
}
