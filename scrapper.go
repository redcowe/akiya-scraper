package main

import (
	"fmt"

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

	// On every a element which has href attribute call callback
	c.OnHTML("section.propety", func(e *colly.HTMLElement) {
		akiyaHTML := e.DOM
		link := e.Attr("href")
		akiya := Akiya{
			Title:  akiyaHTML.Find("div.propetyTitle").Text(),
			Link:   link,
			Price:  akiyaHTML.Find("dl.price").Text(),
			Layout: akiyaHTML.Find("ul.flex").Nodes[0].FirstChild.NextSibling.Data,
		}
		fmt.Println(akiya.Title, akiya.Price, akiya.Link, akiya.Layout)
		akiyaSlice = append(akiyaSlice, akiya)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Start scraping on https://hackerspaces.org
	c.Visit("https://www.akiya-athome.jp/buy/34/?br_kbn=buy&pref_cd=34&page=1&search_sort=kokai_date&item_count=100")
}
