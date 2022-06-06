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

// Helper function for checking if database connection is initalized
func checkInitialized() {
	if !initalized {
		err := connectDB()
		if err != nil {
			fmt.Println("Unabled to connect to server", err)
		}
	}
}

//helper function for creating tables
func CreateTable(i interface{}) error {
	checkInitialized()
	err := DB.AutoMigrate(&i)
	return err
}

// Helper function to insert a new Akiya
func InsertAkiyaBuy(a *akiya.Akiya) {
	checkInitialized()
	DB.Create(&a)
}

//Helper function to get all Akiyas in DB
func GetAkiyasBuy() akiya.Akiyas {
	checkInitialized()
	var akiyas []akiya.Akiya
	DB.Find(&akiyas)
	return akiyas
}

func ClearDBBuy() {
	checkInitialized()
	fmt.Println("Clearing Buy DB...")
	//Reseting primary ID count and clearning DB
	DB.Exec("TRUNCATE akiyas RESTART IDENTITY;")
}

func InsertAkiyaRent(a *akiya.AkiyaRent) {
	DB.Create(*a)
}

func GetAkiyaRent() akiya.AkiyasRent {
	checkInitialized()
	var akiyasRent []akiya.AkiyaRent
	DB.Find(&akiyasRent)
	return akiyasRent
}

func ClearDBRent() {
	checkInitialized()
	fmt.Println("Clearing Rent DB...")
	DB.Exec("TRUNCATE akiya_rents RESTART IDENTITY;")
}
