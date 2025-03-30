package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

  "github.com/GnotAI/skilltrade/internal/users"
)

var DB *gorm.DB

func init() {
  connectDB()
}

// ConnectDB initializes the database connection
func connectDB() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=require TimeZone=UTC",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
    DisableAutomaticPing: true, // Prevents repeated CURRENT_DATABASE() checks
    PrepareStmt: true, // Caches prepared statements to speed up queries
    SkipDefaultTransaction: true, // Avoids extra overhead
	})
	if err != nil {
		log.Fatal("Failed to connect to the database: ", err)
	}

  if (os.Getenv("APP_ENV") == "development") {
		db.Logger = logger.Default.LogMode(logger.Info)
  } else {
		db.Logger = logger.Default.LogMode(logger.Silent)
  }


	// Get the underlying SQL database connection
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Failed to get database instance: ", err)
	}

	// Set connection pool settings
	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)

	DB = db
	log.Println("Database connected successfully")


	// Run migrations
	if err := users.MigrateUsersTable(db); err != nil {
		log.Fatal("Failed to migrate users table:", err)
	}
}

// DisconnectDB gracefully closes the database connection
func DisconnectDB() {
	if DB != nil {
		sqlDB, err := DB.DB()
		if err != nil {
			log.Println("Error getting database instance for closing: ", err)
			return
		}
		if err := sqlDB.Close(); err != nil {
			log.Println("Error closing database connection: ", err)
		} else {
			log.Println("Database connection closed successfully")
		}
	}
}
