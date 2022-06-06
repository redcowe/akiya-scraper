package database

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/redcowe/akiya-scrapper/akiya"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//global DB and initalized status variables
var DB, initalized = &gorm.DB{}, false

//setting up connection
func connectDB() error {
	//Setting env variables and connection string
	_ = godotenv.Load(".env")
	DB_HOST, DB_USER, DB_PASSWORD := os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD")
	DB_NAME, DB_PORT := os.Getenv("DB_NAME"), os.Getenv("DB_PORT")
	dsn := "host=" + DB_HOST + " user=" + DB_USER + " password=" + DB_PASSWORD + " dbname=" + DB_NAME + " port=" + DB_PORT + " sslmode=require"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Printf("Unable to connect to databse, %v\n", err)
		os.Exit(1)
		initalized = false
		return err
	}

	DB = db
	initalized = true
	return nil
}

//helper function for creating tables
func CreateTable(i interface{}) error {
	if !initalized {
		err := connectDB()
		if err != nil {
			fmt.Println("Unabled to connect to server", err)
		}
	}
	err := DB.AutoMigrate(&i)
	return err
}

// Helper function to insert a new Akiya
func InsertAkiyaBuy(a *akiya.Akiya) {

	if !initalized {
		err := connectDB()
		if err != nil {
			fmt.Println("Unabled to connect to server", err)
		}
	}
	DB.Create(&a)
}

//Helper function to get all Akiyas in DB
func GetAkiyasBuy() akiya.Akiyas {
	if !initalized {
		err := connectDB()
		if err != nil {
			fmt.Println("Unabled to connect to server", err)
		}
	}
	var akiyas []akiya.Akiya
	DB.Find(&akiyas)
	return akiyas
}

func ClearDBBuy() {
	if !initalized {
		err := connectDB()
		if err != nil {
			fmt.Println("Unabled to connect to server", err)
		}
	}
	fmt.Println("Clearing DB...")
	//Reseting primary ID count
	DB.Exec("TRUNCATE akiyas RESTART IDENTITY;")
}
