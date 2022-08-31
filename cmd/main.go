package main

import (
	"gitlab.com/ricardo134/auth-service/boot"
)

func main() {
	boot.LoadEnv()
	boot.LoadDb()
	boot.LoadServices()

	boot.ServeHTTP()
}
