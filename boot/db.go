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
		fmt.Sprint("postgres://", dbUser, ":", dbPassword, "@", dbHost, ":", dbPort, "?sslmode=disable")), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "ricardo.",
		},
	})
	if err != nil {
		errors.CannotConnectToDb(dbHost, dbPort)
	}

	_ = client.Migrator().DropTable(entities.User{})
	err = client.AutoMigrate(&entities.User{})
	if err != nil {
		log.Fatal("could not migrate db, exiting...")
	}

	if debug {
		client.Create(&entities.User{
			Username: "test_user",
			Password: "test_password",
		})
	}
}
