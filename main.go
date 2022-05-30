package main

import (
	"github.com/redcowe/akiya-scrapper/scrapper"
)

func main() {
	locationID := "03"
	url := "https://www.akiya-athome.jp/buy/" + locationID + "/" + "?br_kbn=buy&pref_cd=" + locationID + "&page=1&search_sort=kokai_date&item_count=500"
	scrapper.ScrapeAkiyas(url, locationID)

}
