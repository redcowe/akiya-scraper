package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/briandowns/spinner"
	"github.com/redcowe/akiya-scrapper/database"
	"github.com/redcowe/akiya-scrapper/scrapper"
)

func main() {

	database.ClearDBRent()
	database.ClearDBBuy()
	s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
	s.Start()

	//Getting content from every page
	for i := 1; i <= 47; i++ {
		locationID := strconv.Itoa(i)
		if i < 10 {
			locationID = "0" + strconv.Itoa(i)
		}
		scrapper.ScrapeAkiyasBuy(locationID)
		scrapper.ScrapeAkiyasRent(locationID)
		fmt.Println("----------------------------------------------------------------------------------")
	}

	// //Displaying IDs
	Akiyas := database.GetAkiyaRent()
	for _, akiya := range Akiyas {
		fmt.Println(akiya.ID, akiya.Link)
		fmt.Println("----------")
	}
	s.Stop()
}
