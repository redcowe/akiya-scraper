package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/gocolly/colly"
	"github.com/joho/godotenv"
)

//Akiya structure declaration
type Akiya struct {
	Title      string `json:"title"`
	Link       string `json:"link"`
	Price      string `json:"price"`
	Desc       string `json:"desc"`
	Area       string `json:"area"`
	Type       string `json:"type"`
	Location   string `json:"location"`
	LocationID string `json:"locationId"`
}

//helper function for converting empty descriptions into proper format
func descEmptyConvert(s *string) string {
	if *s == "" {
		*s = "N/A"
	}
	return *s
}

//function for JSONifying the object and writing to a file
func writeFile(data []Akiya, locationID string) {
	file, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return
	}
	_ = ioutil.WriteFile(locationID+".json", file, 0644)
}
func main() {

	//Setting ENV variables
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Failed to load .env file")
		os.Exit(1)
	}

	//Location and URL that can be edited
	locationID := "02"
	url := "https://www.akiya-athome.jp/buy/" + locationID + "/" + "?br_kbn=buy&pref_cd=" + locationID + "&page=1&search_sort=kokai_date&item_count=10"
	//https://www.akiya-athome.jp/buy/38/?br_kbn=buy&pref_cd=38&page=1&search_sort=kokai_date&item_count=50
	//https://www.akiya-athome.jp/buy/" + locationID + "/" + "?br_kbn=buy&pref_cd=" + locationID + "&page=1&search_sort=kokai_date&item_count=100

	akiyaSlice := []Akiya{}
	c := colly.NewCollector(
		//Setting domains
		colly.AllowedDomains(
			"akiya-athome.jp/",
			"https://www.akiya-athome.jp",
			"www.akiya-athome.jp",
		),
	)

	c.OnHTML("section.propety", func(e *colly.HTMLElement) {
		akiyaHTML := e.DOM
		desc := akiyaHTML.Find("div.description").Text()
		akiya := Akiya{ //Filling akiya object
			Title:      strings.TrimSpace(akiyaHTML.Find("div.propetyTitle").Find("a").Text()),
			Link:       akiyaHTML.Find("div.propetyTitle").Find("a").AttrOr("href", "N/A"),
			Price:      akiyaHTML.Find("dl.price").Find("dd").Text(),
			Desc:       descEmptyConvert(&desc),
			Area:       akiyaHTML.Find("ul.flex").Find("li").Find("dl").Find("dd:contains(㎡)").Text(),
			Type:       akiyaHTML.Find("div.objectTitle.cf").Find("span.objectCategory.objectCategory_buy").Text(),
			Location:   akiyaHTML.Find("ul.all").Find("li").Find("dt:contains(所在地)").Next().Text(),
			LocationID: locationID,
		}
		akiyaJSON, err := json.MarshalIndent(akiya, "", " ")
		if err != nil {
			return
		}
		fmt.Println(string(akiyaJSON))
		akiyaSlice = append(akiyaSlice, akiya)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.Visit(url)
	writeFile(akiyaSlice, locationID)
}
