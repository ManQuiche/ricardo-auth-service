package boot

import (
	"auth-service/pkg/errors"
	"os"
)

var (
	dbHost        string
	dbPort        string
	dbUser        string
	dbPassword    string
	port          string
	url           string
	accessSecret  string
	refreshSecret string
)

func LoadEnv() {
	dbHost = env("DB_HOST")
	dbPort = env("DB_PORT")
	dbUser = env("DB_USER")
	dbPassword = env("DB_PASSWORD")
	port = env("PORT")
	url = env("URL")
	accessSecret = env("ACCESS_SECRET")
	refreshSecret = env("REFRESH_SECRET")
}

func env(name string) string {
	str, ok := os.LookupEnv(name)
	if !ok {
		errors.MissingEnvVarF(name)
	}

	return str
}
