package boot

import (
	"auth-service/pkg/errors"
	"os"
)

var (
	dbHost     string
	dbUrl      string
	dbUser     string
	dbPassword string
	port       string
	url        string
)

func LoadEnv() {
	dbHost = env("DB_HOST")
	dbUrl = env("DB_URL")
	dbUser = env("DB_USER")
	dbPassword = env("DB_PASSWORD")
	port = env("PORT")
	url = env("URL")
}

func env(name string) string {
	str, ok := os.LookupEnv(name)
	if !ok {
		errors.MissingEnvVarF(name)
	}

	return str
}
