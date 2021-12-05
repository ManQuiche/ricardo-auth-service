package boot

import (
	"auth-service/pkg/errors"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
)

var (
	client *sql.DB
)

func LoadDb() {

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	// TODO: finish that shiiit
	//connectionString := fmt.Sprintf("")
	client, err := sql.Open("cockroach", fmt.Sprintf(user, ":", password, "@", host, ":", port))
	if err != nil {
		errors.CannotConnectToDb(host, port)
	}

	err = client.Ping()
	if err != nil {
		errors.CannotConnectToDb(host, port)
	}
}
