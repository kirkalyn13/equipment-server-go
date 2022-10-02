package config

import (
	"log"
	"os"

	"database/sql"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	log.Print("Connecting to database...")
	var err error

	dsn := os.Getenv("DB_URL")
	sqlDB, err := sql.Open("mysql", dsn)
	DB, err = gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database.")
	}
	log.Print("Successfully connected to the database.")
}
