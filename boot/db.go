package boot

import (
	"auth-service/pkg/errors"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

var (
	client *sql.DB
)

func LoadDb() {
	// TODO: finish that shiiit
	//connectionString := fmt.Sprintf("")
	client, err := sql.Open("cockroach", fmt.Sprintf(dbUser, ":", dbPassword, "@", dbHost, ":", dbPort))
	if err != nil {
		errors.CannotConnectToDb(dbHost, dbPort)
	}

	err = client.Ping()
	if err != nil {
		errors.CannotConnectToDb(dbHost, dbPort)
	}
}
