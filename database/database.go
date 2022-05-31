package database

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/redcowe/akiya-scrapper/akiya"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func connectDB() *gorm.DB {

	_ = godotenv.Load(".env")
	//Setting env variables and connection string
	DB_HOST, DB_USER, DB_PASSWORD := os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD")
	DB_NAME, DB_PORT := os.Getenv("DB_NAME"), os.Getenv("DB_PORT")
	dsn := "host=" + DB_HOST + " user=" + DB_USER + " password=" + DB_PASSWORD + " dbname=" + DB_NAME + " port=" + DB_PORT + " sslmode=disable"

	fmt.Println(DB_HOST)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("Unable to connect to databse, %v\n", err)
		os.Exit(1)
	}

	return db
}

func CreateTable() {
	db := connectDB()
	db.Migrator().CreateTable(&akiya.Akiya{})
}
