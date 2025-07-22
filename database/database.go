package database

import (
	"fmt"
	"log"
	"os"
	"sync"
	"taxi-service/models"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	instance *gorm.DB
	once     sync.Once
)

// GetDB returns the singleton database instance
func GetDB() *gorm.DB {
	once.Do(func() {
		var db *gorm.DB
		var err error

		if os.Getenv("DB_DRIVER") == "sqlite" {
			// Use SQLite for testing
			db, err = gorm.Open(sqlite.Open(os.Getenv("DB_NAME")), &gorm.Config{
				Logger: logger.Default.LogMode(logger.Info),
			})
		} else {
			// Use PostgreSQL for production
			dsn := fmt.Sprintf(
				"host=%s user=%s password=%s dbname=%s port=%s",
				os.Getenv("DB_HOST"),
				os.Getenv("DB_USER"),
				os.Getenv("DB_PASSWORD"),
				os.Getenv("DB_NAME"),
				os.Getenv("DB_PORT"),
			)

			db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
				Logger: logger.Default.LogMode(logger.Info),
			})
		}

		if err != nil {
			log.Fatal("Failed to connect to database. \n", err)
		}

		log.Println("connected")
		db.Logger = logger.Default.LogMode(logger.Info)
		log.Println("running migrations")

		db.AutoMigrate(&models.DummyUser{})
		instance = db
	})
	return instance
}

// ConnectDb initializes the database connection
func ConnectDb() {
	GetDB()
}
