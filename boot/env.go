package boot

import (
	"auth-service/pkg/errors"
	"os"
)

func LoadEnv() {

}

func LoadOneEnv(name string) string {
	str, ok := os.LookupEnv(name)
	if !ok {
		errors.MissingEnvVarF(name)
	}

	return str
}
