package main

import (
	"github.com/redcowe/akiya-scrapper/scrapper"
)

func main() {
	locationID := "08"
	scrapper.ScrapeAkiyas(locationID)
}
