package utils

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	onceDbConnection sync.Once
	db               *gorm.DB
)

func GetDbConnection() *gorm.DB {
	onceDbConnection.Do(func() {

		// Fetch configs from .env
		host := os.Getenv("DB_HOST")
		port := os.Getenv("DB_PORT")
		user := os.Getenv("DB_USER")
		password := os.Getenv("DB_PASSWORD")
		dbName := os.Getenv("DB_NAME")

		// Build DSN
		// parseTime=true => MySQL datetime -> Go time.Time
		// loc=Asia%2FKolkata => correct timezone handling
		dsn := fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Asia%%2FKolkata",
			user, password, host, port, dbName,
		)

		var err error

		// Connect to MySQL
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info), // SQL logs
		})

		if err != nil {
			log.Fatalf("Failed to connect with MySQL: %v", err)
		}

		// ----------------------
		//   Connection Pooling
		// ----------------------
		sqlDB, err := db.DB()
		if err != nil {
			log.Fatalf("Failed to get generic DB: %v", err)
		}

		sqlDB.SetMaxOpenConns(50)               // Maximum active connections
		sqlDB.SetMaxIdleConns(25)               // Idle connections ready for reuse
		sqlDB.SetConnMaxLifetime(2 * time.Hour) // Recycle after 2 hours
		sqlDB.SetConnMaxIdleTime(30 * time.Minute)

		log.Println("MySQL connected successfully!")
	})

	return db
}
