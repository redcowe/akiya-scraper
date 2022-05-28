package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

//Akiya structure declaration
type Akiya struct {
	Title  string `json:"title"`
	Link   string `json:"link"`
	Price  string `json:"price"`
	Layout string `json:"layout"`
	// Area     string `json:"area"`
	// Type     string `json:"type"`
	// Location string `json:location`
}

func main() {

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
		link := e.Attr("href")

		akiya := Akiya{
			Title:  strings.TrimSpace(akiyaHTML.Find("div.propetyTitle").Find("a[href]").Text()),
			Link:   link,
			Price:  akiyaHTML.Find("dl.price").Find("dd").Text(),
			Layout: akiyaHTML.Find("ul.flex").Find("li").Find("dl").Find("dd:contains(DK)").Text(),
		}
		akiyaSlice = append(akiyaSlice, akiya)
		akiyaJSON, err := json.MarshalIndent(akiya, "", " ")
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(akiyaJSON))
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.Visit("https://www.akiya-athome.jp/buy/01/?br_kbn=buy&pref_cd=01&page=1&search_sort=kokai_date&item_count=10")

}
