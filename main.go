package main

import (
	"github.com/redcowe/akiya-scrapper/database"
)

func main() {
	database.ClearDBRent()
	// //Getting content from every page
	// for i := 1; i <= 47; i++ {
	// 	locationID := strconv.Itoa(i)
	// 	if i < 10 {
	// 		locationID = "0" + strconv.Itoa(i)
	// 	}
	// 	scrapper.ScrapeAkiyasBuy(locationID)
	// 	fmt.Println("-------------------------------------------------------------------------------------------------------------------")
	// }

	//Displaying IDs
	// Akiyas := database.GetAkiyasBuy()
	// for _, akiya := range Akiyas {
	// 	fmt.Println(akiya.ID)
	// 	fmt.Println("----------")
	// }
}
