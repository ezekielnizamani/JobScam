package database

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() error {
	var err error

	dsn := os.Getenv("DATABASE_URL") // or use a direct connection string like below
	// dsn := "host=localhost user=username password=password dbname=jobscam_db port=5432 sslmode=disable TimeZone=Asia/Karachi"

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	log.Println("Connected to PostgreSQL with GORM!")
	return nil
}

func AutoMigrate(models ...interface{}) error {
	return DB.AutoMigrate(models...)
}

func GetDB() *gorm.DB {
	return DB
}
