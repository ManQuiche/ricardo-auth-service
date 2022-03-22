package main

import (
	"ricardo/auth-service/boot"
)

func main() {
	boot.LoadEnv()
	boot.LoadDb()
	boot.LoadServices()

	boot.ServeHTTP()
}
