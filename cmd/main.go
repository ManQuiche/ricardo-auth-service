package main

import (
	"ricardo/auth-service/boot"
	"ricardo/auth-service/internal/driven/firebase"
)

func main() {
	boot.LoadEnv()
	boot.LoadDb()
	boot.LoadServices()

	firebase.InitFirebaseSDK()

	boot.ServeHTTP()
}
