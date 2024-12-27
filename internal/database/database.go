package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB global variable to hold the GORM DB instance
var db *gorm.DB

// Connect initializes the database connection using GORM
func Connect(host, port, user, password, dbname string) {
	// Connection string
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// Open a GORM connection
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}
}

// GetDB returns the GORM DB instance
func GetDB() *gorm.DB {
	return db
}
