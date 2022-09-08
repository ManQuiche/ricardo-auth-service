package boot

import (
	"context"
	"fmt"
	_ "github.com/lib/pq"
	errors2 "github.com/pkg/errors"
	"gitlab.com/ricardo134/auth-service/internal/core/entities"
	"gitlab.com/ricardo134/auth-service/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
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
}

func LoadAdditionalData() {
	if authenticateService == nil {
		log.Fatal("cannot load additional data: authenticateService is nil (is it run after LoadServices ?)")
	}

	if debug {
		err := authenticateService.Save(context.Background(), entities.User{
			Username: "test_user",
			Password: "test_password",
			Email:    "test@test.fr",
		})

		if err != nil {
			log.Fatal(errors2.Wrap(err, "could not save debug user"))
		}
	}
}
