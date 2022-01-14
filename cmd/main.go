package main

import (
	"auth-service/boot"
	"auth-service/internal/driven/firebase"
)

func main() {
	boot.LoadEnv()
	boot.LoadDb()
	boot.LoadServices()

	firebase.InitFirebaseSDK()

	boot.ServeHTTP()
}
