package main

import (
	//"github.com/redcowe/akiya-scrapper/scrapper"
	"fmt"

	"github.com/redcowe/akiya-scrapper/database"
)

func main() {

	//locationID := "44"
	//scrapper.ScrapeAkiyas(locationID)
	Akiyas := database.GetAkiyas()
	for _, akiya := range Akiyas {
		fmt.Println(akiya.LocationID)
		fmt.Println("----------")
	}
}
