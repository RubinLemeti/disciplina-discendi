package config

import (
	// "database/sql"
	"log"
	// "time"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

// InitGORM initializes the database connection pool
func InitGORM() *gorm.DB {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	database := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	sslMode := os.Getenv("SSL_MODE")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		host, user, password, database, port, sslMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect database: %v", err)
	}

	// Get the underlying sql.DB instance for connection pooling settings
	// sqlDB, err := db.DB()
	// if err != nil {
	// 	log.Fatalf("Failed to get sql.DB: %v", err)
	// }

	// // Configure connection pooling
	// sqlDB.SetMaxOpenConns(10)         // Maximum open connections
	// sqlDB.SetMaxIdleConns(2)          // Minimum idle connections
	// sqlDB.SetConnMaxLifetime(30 * time.Minute) // Connection lifetime

	log.Println("Connected to PostgreSQL with GORM (without models)")
	return db
}
