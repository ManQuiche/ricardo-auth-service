package main

import (
	"gitlab.com/ricardo134/auth-service/boot"
)

func main() {
	boot.LoadEnv()
	boot.LoadDb()
	boot.InitFirebaseApp()
	boot.LoadServices()
	boot.LoadTracing()
	//boot.LoadAdditionalData()

	boot.ServeHTTP()
}
