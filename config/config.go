package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	L "rest-api-implementation/middleware/logger"
	"rest-api-implementation/models"
)

// InitDB initializes the database connection
func InitDB() (*gorm.DB, error) {
	// Open the database connection
	db, err := OpenDB()
	if err != nil {
		return nil, err
	}

	// Auto-migrate schema
	err = db.AutoMigrate(&models.User{}, &models.Token{})
	if err != nil {
		L.RaiLog("E", "Error Occurred During Auto Migration", err)
		return nil, err
	}
	L.RaiLog("D", "DB Models Successful", nil)
	defer CloseDB(db)
	return db, nil
}

var SecretKey string

// OpenDB opens the database connection
func OpenDB() (*gorm.DB, error) {
	// Load .env file
	err := godotenv.Load("./infrastructure/.env")
	if err != nil {
		L.RaiLog("E", "Error Loading .env file", err)
	}
	// Connection string
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Kolkata",
		"db",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
		"5432",
	)
	// Get the secret key from the environment variables
	SecretKey = os.Getenv("SECRET")

	// Connect to PostgresSQL database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		L.RaiLog("E", "Error Occurred while opening dbConnection", err)
		return nil, err
	}
	//get secret key from .env file

	return db, nil
}

// CloseDB closes the database connection
func CloseDB(db *gorm.DB) {
	if db == nil {
		L.RaiLog("D", "Database connection closed", nil)
		return
	}

	sqlDB, err := db.DB()
	if err != nil {
		L.RaiLog("E", "Error getting database instance", err)
		return
	}

	err = sqlDB.Close()
	if err != nil {
		L.RaiLog("E", "Error closing database", err)
		return
	}

	L.RaiLog("D", "Database connection closed", nil)
}
