package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectDatabse() {
	connectionString := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"

	DB, err = gorm.Open(postgres.Open(connectionString))
	if err != nil {
		log.Panic("Failed to connect to database!")
	}
}
