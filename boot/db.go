package boot

import (
	"fmt"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
	"ricardo/auth-service/internal/core/entities"
	"ricardo/auth-service/pkg/errors"
)

var (
	client *gorm.DB
)

func LoadDb() {
	// TODO: finish that shiiit
	//connectionString := fmt.Sprintf("")

	var err error
	client, err = gorm.Open(postgres.Open(
		fmt.Sprint("postgres://", dbUser, ":", dbPassword, "@", dbHost, ":", dbPort, "/", dbDatabase, "?sslmode=disable")), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: fmt.Sprint(dbSchema, "."),
		},
	})
	if err != nil {
		errors.CannotConnectToDb(dbHost, dbPort)
	}

	err = client.AutoMigrate(&entities.User{})
	if err != nil {
		log.Fatal("could not migrate db, exiting...")
	}

	if debug {
		var user entities.User
		client.FirstOrCreate(&user, entities.User{
			Username: "test_user",
			Password: "test_password",
			Email:    "test@test.fr",
		})
	}
}
