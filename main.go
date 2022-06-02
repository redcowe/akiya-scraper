package main

import (
	"fmt"
	"strconv"

	"github.com/redcowe/akiya-scrapper/database"
	"github.com/redcowe/akiya-scrapper/scrapper"
)

func main() {

	//clearing db before updated new values
	database.ClearDB()

	//Getting content from every page
	for i := 1; i <= 47; i++ {
		locationID := strconv.Itoa(i)
		if i < 10 {
			locationID = "0" + strconv.Itoa(i)
		}
		scrapper.ScrapeAkiyas(locationID)
		fmt.Println("-------------------------------------------------------------------------------------------------------------------")
	}

	//Displaying IDs
	Akiyas := database.GetAkiyas()
	for _, akiya := range Akiyas {
		fmt.Println(akiya.ID)
		fmt.Println("----------")
	}
}
