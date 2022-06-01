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

	for i := 10; i <= 12; i++ {
		scrapper.ScrapeAkiyas(strconv.Itoa(i))
		fmt.Println("------------")
	}
	Akiyas := database.GetAkiyas()
	for _, akiya := range Akiyas {
		fmt.Println(akiya.ID)
		fmt.Println("----------")
	}
}
