package scrapper

import (
	"fmt"
	"strings"
	"time"

	"github.com/gocolly/colly"
	"github.com/redcowe/akiya-scrapper/akiya"
	"github.com/redcowe/akiya-scrapper/database"
)

//Variables for ID and collector

//helper function for converting empty descriptions into proper format
func convertEmptyDesc(s *string) string {
	if *s == "" {
		*s = "N/A"
	}
	return *s
}

func ScrapeAkiyasBuy(locationID string) {
	c := colly.NewCollector(
		//Setting domains
		colly.AllowedDomains(
			"akiya-athome.jp/",
			"https://www.akiya-athome.jp",
			"www.akiya-athome.jp",
		),
	)

	url := "https://www.akiya-athome.jp/buy/" + locationID + "/" + "?br_kbn=buy&pref_cd=" + locationID + "&page=1&search_sort=kokai_date&item_count=500"

	//Increasing timeout for larger requests
	c.SetRequestTimeout(time.Duration(35) * time.Second)

	c.OnHTML("section.propety", func(e *colly.HTMLElement) {
		akiyaHTML := e.DOM
		desc := akiyaHTML.Find("div.description").Text()
		//Filling akiya object
		akiya := akiya.Akiya{
			Title:      strings.TrimSpace(akiyaHTML.Find("div.propetyTitle").Find("a").Text()),
			Link:       akiyaHTML.Find("div.propetyTitle").Find("a").AttrOr("href", "N/A"),
			Price:      akiyaHTML.Find("dl.price").Find("dd").Text(),
			Desc:       convertEmptyDesc(&desc),
			Area:       akiyaHTML.Find("ul.flex").Find("li").Find("dl").Find("dd:contains(㎡)").Text(),
			Type:       akiyaHTML.Find("div.objectTitle.cf").Find("span.objectCategory.objectCategory_buy").Text(),
			Location:   akiyaHTML.Find("ul.all").Find("li").Find("dt:contains(所在地)").Next().Text(),
			LocationID: locationID,
		}
		database.InsertAkiyaBuy(&akiya)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
	})

	c.Visit(url)
}

func ScrapeAkiyasRent(locationId string) {

	c := colly.NewCollector(
		//Setting domains
		colly.AllowedDomains(
			"akiya-athome.jp/",
			"https://www.akiya-athome.jp",
			"www.akiya-athome.jp",
		),
	)

	url := "https://www.akiya-athome.jp/rent/" + locationId + "/" + "?br_kbn=buy&pref_cd=" + locationId + "&page=1&search_sort=kokai_date&item_count=500"

	c.SetRequestTimeout(time.Duration(20) * time.Second)

	c.OnHTML("section.propety", func(e *colly.HTMLElement) {
		akiyaHTML := e.DOM
		desc := akiyaHTML.Find("div.description").Text()

		akiya := akiya.AkiyaRent{
			Title:      strings.TrimSpace(akiyaHTML.Find("div.propetyTitle").Find("a").Text()),
			Link:       akiyaHTML.Find("div.propetyTitle").Find("a").AttrOr("href", "N/A"),
			Rent:       akiyaHTML.Find("dl.price").Find("dd").Text(),
			Desc:       convertEmptyDesc(&desc),
			Area:       akiyaHTML.Find("ul.flex").Find("li").Find("dl").Find("dd:contains(㎡)").Text(),
			Location:   akiyaHTML.Find("ul.all").Find("li").Find("dt:contains(所在地)").Next().Text(),
			LocationID: locationId,
			Layout:     akiyaHTML.Find("ul.flex").Find("li").Find("dt:contains(間取)").Next().Text(),
			WhenBuilt:  akiyaHTML.Find("ul.all").Find("li").Find("dt:contains(築年月)").Next().Text(),
			Type:       akiyaHTML.Find("ul.flex").Find("li").Find("dt:contains(物件種目)").Next().Text(),
		}
		database.InsertAkiyaRent(&akiya)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
	})

	c.Visit(url)
}
