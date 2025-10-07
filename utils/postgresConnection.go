package utils

import (
	"fmt"
	"log"
	"os"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	onceDbConnection sync.Once
	db               *gorm.DB
)

func GetDbConnection() *gorm.DB {
	onceDbConnection.Do(func() {
		// fetching data from .env file

		host := os.Getenv("DB_HOST")
		port := os.Getenv("DB_PORT")
		user := os.Getenv("DB_USER")
		password := os.Getenv("DB_PASSWORD")
		dbName := os.Getenv("DB_NAME")

		// making connection with db
		dsn := fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Kolkata",
			host, port, user, password, dbName,
		)

		var err error

		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

		if err != nil {
			log.Fatalf("Failed to connect with DB: %v", err)
		}
	})
	return db
}
