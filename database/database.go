package database

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func DatabaseInit() {
	errLoad := godotenv.Load()
	if errLoad != nil {
		panic("Failed to load ENV")
	}

	var err error
	name := os.Getenv("DATABASE_NAME")
	username := os.Getenv("DATABASE_USERNAME")
	password := os.Getenv("DATABASE_PASSWORD")
	url := os.Getenv("DATABASE_URL")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable", url, username, password, name)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("-- CANT CONNECT TO DATABASE ! --")
	}

	fmt.Println("-- CONNECTED TO DATABASE --")
}
