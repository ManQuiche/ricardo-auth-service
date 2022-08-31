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
	dbSchema      string
	port          string
	refreshSecret string
	url           string

	natsURL           string
	natsUsr           string
	natsPwd           string
	natsRegisterTopic string

	noFirebase bool
	debug      bool
)

func LoadEnv() {
	accessSecret = env("ACCESS_SECRET")
	dbHost = env("DB_HOST")
	dbPassword = env("DB_PASSWORD")
	dbPort = env("DB_PORT")
	dbUser = env("DB_USER")
	dbDatabase = env("DB_DATABASE")
	dbSchema = env("DB_SCHEMA")
	debug = envBool("DEBUG")
	noFirebase = envBool("NO_FIREBASE")
	port = env("PORT")
	refreshSecret = env("REFRESH_SECRET")
	url = env("URL")

	natsURL = env("NATS_URL")
	natsUsr = env("NATS_USR")
	natsPwd = env("NATS_PWD")
	natsRegisterTopic = env("NATS_REGISTER_TOPIC")
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
