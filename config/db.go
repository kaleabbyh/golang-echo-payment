package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


func ConnectDB() (*gorm.DB, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	dbHost     := os.Getenv("DB_HOST")
	dbPort 	   := os.Getenv("DB_PORT")
	dbUser	   := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName	   := os.Getenv("DB_NAME")
	// dbUrl	   := os.Getenv("DB_URL")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)
	fmt.Println(dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
		log.Fatal("Error  creating user table:", err)
		return nil, err
		
	}
	fmt.Println("DB connected successfully")
	
	return db, nil
}