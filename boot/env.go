package boot

import (
	"gitlab.com/ricardo134/auth-service/pkg/errors"
	"log"
	"os"
	"strconv"
)

var (
	accessSecret  string
	dbHost        string
	dbPassword    string
	dbPort        string
	dbUser        string
	dbDatabase    string
	port          string
	refreshSecret string
	url           string

	natsURL         string
	natsUsr         string
	natsPwd         string
	natsUserCreated string
	natsUserUpdated string
	natsUserDeleted string

	noFirebase bool
	debug      bool

	environment     string
	tracingEndpoint string
)

func LoadEnv() {
	accessSecret = env("ACCESS_SECRET")
	refreshSecret = env("REFRESH_SECRET")

	dbHost = env("DB_HOST")
	dbPassword = env("DB_PASSWORD")
	dbPort = env("DB_PORT")
	dbUser = env("DB_USER")
	dbDatabase = env("DB_DATABASE")

	debug = envBool("DEBUG")
	noFirebase = envBool("NO_FIREBASE")
	port = env("PORT")
	url = env("URL")

	natsURL = env("NATS_URL")
	natsUsr = env("NATS_USR")
	natsPwd = env("NATS_PWD")
	natsUserCreated = env("NATS_USER_CREATED")
	natsUserUpdated = env("NATS_USER_UPDATED")
	natsUserDeleted = env("NATS_USER_DELETED")

	// DMZ, QA, PROD ?
	environment = env("ENVIRONMENT")

	tracingEndpoint = env("TRACING_ENDPOINT")
}

func envBool(name string) bool {
	res, err := strconv.ParseBool(env(name))
	if err != nil {
		log.Fatalf("env var %s needs to be of boolean type", name)
	}

	return res
}

func env(name string) string {
	str, ok := os.LookupEnv(name)
	if !ok {
		errors.MissingEnvVarF(name)
	}

	return str
}
