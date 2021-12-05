package boot

import (
	"auth-service/pkg/errors"
	"os"
)

func LoadEnv() {
	LoadOneEnv("DB_HOST")
	LoadOneEnv("DB_HOST")
	LoadOneEnv("PORT")
	LoadOneEnv("URL")
}

func LoadOneEnv(name string) string {
	str, ok := os.LookupEnv(name)
	if !ok {
		errors.MissingEnvVarF(name)
	}

	return str
}
