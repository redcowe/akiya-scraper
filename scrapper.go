package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

//Akiya structure declaration
type Akiya struct {
	Title string `json:"title"`
	//Link	  string `json:"link"`
	//Price     string `json:"price"`
	//Layout    string `json:"layout"`
	//Area      string `json:"area"`
	//Type      string `json:"type"`
	//ConYear   string `json:conyear"`
	//Location  string `json:location"`
	//Transport string `json:transport"`
}

func main() {

	akiyaSlice := []Akiya{}
	c := colly.NewCollector(
		// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
		colly.AllowedDomains("akiya-athome.jp/"),
	)

	// On every a element which has href attribute call callback
	c.OnHTML("section.propety", func(e *colly.HTMLElement) {
		akiyaHTML := e.DOM

		akiya := Akiya{
			Title: akiyaHTML.Find("div.propetyTitle").Text(),
		}
		fmt.Println(akiya.Title)
		akiyaSlice = append(akiyaSlice, akiya)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})
	// Start scraping on https://hackerspaces.org
	c.Visit("https://www.akiya-athome.jp/buy/34/")
}
