package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// The function to connect database (Postgres)
func OpenDB(dsn string) *gorm.DB {
	// Connect the database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	return db
}

// The function to close database connection
func CloseDB(db *gorm.DB) {
	// Close the database connection
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to close database connection")
	}

	dbSQL.Close()
}
