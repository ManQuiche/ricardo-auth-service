package boot

import (
	"log"
	"os"
	"ricardo/auth-service/pkg/errors"
	"strconv"
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

	debug bool
)

func LoadEnv() {
	var err error

	dbHost = env("DB_HOST")
	dbPort = env("DB_PORT")
	dbUser = env("DB_USER")
	dbPassword = env("DB_PASSWORD")
	port = env("PORT")
	url = env("URL")
	accessSecret = env("ACCESS_SECRET")
	refreshSecret = env("REFRESH_SECRET")
	debug, err = strconv.ParseBool(env("DEBUG"))

	if err != nil {
		log.Fatal("env var DEBUG needs to be of boolean type")
	}
}

func env(name string) string {
	str, ok := os.LookupEnv(name)
	if !ok {
		errors.MissingEnvVarF(name)
	}

	return str
}
