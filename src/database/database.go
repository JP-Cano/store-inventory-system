package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"store-inventory-management/src/entities"
)

func New(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("Error connecting to database: %v", err)
		return nil, err
	}
	log.Println("Database connected successfully")
	return db, nil
}

func Close(db *gorm.DB) {
	conn, err := db.DB()
	if err != nil {
		log.Printf("Error retrieving database connection: %v", err)
		return
	}
	err = conn.Close()
	if err != nil {
		log.Printf("Error closing database connection: %v", err)
		return
	}
	log.Println("Database connection closed successfully")
	return
}

func Migrate(db *gorm.DB) {
	createUUIDExtension(db)
	err := db.AutoMigrate(entities.Product{}, entities.Provider{}, entities.Sale{})
	if err != nil {
		log.Printf("Error migrating database: %v", err)
		return
	}
}

func createUUIDExtension(db *gorm.DB) {
	if err := db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`).Error; err != nil {
		log.Fatalf("Failed to create uuid extension: %v", err)
	}
}
