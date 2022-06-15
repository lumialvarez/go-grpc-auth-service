package postgresql

import (
	"github.com/lumialvarez/go-grpc-auth-service/src/infrastructure/platform/dao"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Client struct {
	DB *gorm.DB
}

func Init(url string) Client {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(dao.User{})

	return Client{db}
}
