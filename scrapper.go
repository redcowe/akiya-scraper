package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

//Akiya structure declaration
type Akiya struct {
	Title    string `json:"title"`
	Link     string `json:"link"`
	Price    string `json:"price"`
	Desc     string `json:"desc"`
	Area     string `json:"area"`
	Type     string `json:"type"`
	Location string `json:"location"`
}

//helper function for converting empty descriptions into proper format
func descEmptyConvert(s *string) string {
	if *s == "" {
		*s = "N/A"
	}
	return *s
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
		desc := akiyaHTML.Find("div.description").Text()
		akiya := Akiya{
			Title:    strings.TrimSpace(akiyaHTML.Find("div.propetyTitle").Find("a").Text()),
			Link:     akiyaHTML.Find("div.propetyTitle").Find("a").AttrOr("href", "N/A"),
			Price:    akiyaHTML.Find("dl.price").Find("dd").Text(),
			Desc:     descEmptyConvert(&desc),
			Area:     akiyaHTML.Find("ul.flex").Find("li").Find("dl").Find("dd:contains(㎡)").Text(),
			Type:     akiyaHTML.Find("div.objectTitle.cf").Find("span.objectCategory.objectCategory_buy").Text(),
			Location: akiyaHTML.Find("ul.all").Find("li").Find("dt:contains(所在地)").Next().Text(),
		}
		akiyaSlice = append(akiyaSlice, akiya)

		//JSONifying the object
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

	c.Visit("https://www.akiya-athome.jp/buy/03/?br_kbn=buy&pref_cd=03&page=1&search_sort=kokai_date&item_count=10")

}
