package main

import "github.com/redcowe/akiya-scrapper/scrapper"

func main() {

	locationID := "21"
	scrapper.ScrapeAkiyas(locationID)

	//Akiyas := database.GetAkiyas()
	//for _, akiya := range Akiyas {
	//	fmt.Println(akiya.ID)
	//	fmt.Println("----------")
	//}
}
