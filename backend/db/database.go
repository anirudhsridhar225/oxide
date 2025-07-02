package db

import (
	"fmt"
	"log"
	"os"

	"oxide/models" // Import your models

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Init() {
	var err error
	dsn := os.Getenv("DATABASE_DSN")

	if dsn == "" {
		dsn = "host=localhost user=anirudh password=@nirudh225 dbname=oxide_db port=5432 sslmode=disable"
	}

	// Configure GORM with logger for better debugging
	config := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}

	DB, err = gorm.Open(postgres.Open(dsn), config)

	if err != nil {
		log.Fatal("error opening database. try again\n")
	}

	// Test the connection
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatal("failed to get underlying sql.DB")
	}

	err = sqlDB.Ping()
	if err != nil {
		log.Fatal("failed to ping database")
	}

	fmt.Println("Database connection established successfully")
	fmt.Println("running database migrations")

	migrationErr := DB.AutoMigrate(
		&models.User{},
		&models.Session{},
	)

	if migrationErr != nil {
		log.Fatal("Failed to migrate database:", migrationErr)
	}

	fmt.Println("Database initialization and migration completed successfully!")
}

// Migrate runs the auto-migration for all models
func Migrate() {
	fmt.Println("Running database migrations...")

	err := DB.AutoMigrate(
		&models.User{},
		&models.Session{},
		// Add more models here as you create them
	)

	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	fmt.Println("Database migration completed successfully!")
}
