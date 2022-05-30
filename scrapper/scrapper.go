package scrapper

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"github.com/gocolly/colly"
	"github.com/redcowe/akiya-scrapper/akiya"
)

//helper function for converting empty descriptions into proper format
func convertEmptyDesc(s *string) string {
	if *s == "" {
		*s = "N/A"
	}
	return *s
}

//function for JSONifying the object and writing to a file
func writeFile(data []akiya.Akiya, id string) {
	file, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return
	}
	_ = ioutil.WriteFile(id+".json", file, 0644)
}

func ScrapeAkiyas(locationID string) {

	id := locationID
	url := "https://www.akiya-athome.jp/buy/" + id + "/" + "?br_kbn=buy&pref_cd=" + id + "&page=1&search_sort=kokai_date&item_count=500"

	akiyaSlice := []akiya.Akiya{}
	c := colly.NewCollector(
		//Setting domains
		colly.AllowedDomains(
			"akiya-athome.jp/",
			"https://www.akiya-athome.jp",
			"www.akiya-athome.jp",
		),
	)

	//Increasing timeout for larger requests
	c.SetRequestTimeout(time.Duration(30) * time.Second)

	c.OnHTML("section.propety", func(e *colly.HTMLElement) {
		akiyaHTML := e.DOM
		desc := akiyaHTML.Find("div.description").Text()
		akiya := akiya.Akiya{ //Filling akiya object
			Title:      strings.TrimSpace(akiyaHTML.Find("div.propetyTitle").Find("a").Text()),
			Link:       akiyaHTML.Find("div.propetyTitle").Find("a").AttrOr("href", "N/A"),
			Price:      akiyaHTML.Find("dl.price").Find("dd").Text(),
			Desc:       convertEmptyDesc(&desc),
			Area:       akiyaHTML.Find("ul.flex").Find("li").Find("dl").Find("dd:contains(㎡)").Text(),
			Type:       akiyaHTML.Find("div.objectTitle.cf").Find("span.objectCategory.objectCategory_buy").Text(),
			Location:   akiyaHTML.Find("ul.all").Find("li").Find("dt:contains(所在地)").Next().Text(),
			LocationID: id,
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
	writeFile(akiyaSlice, id)
}
