package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

type Flat struct {
	Title       string
	Description string
	District    string
	Street      string
}

func main() {
	c := colly.NewCollector(
		// Visit only these domains
		colly.AllowedDomains("realty.yandex.ru"),
		// Cache responses to prevent multiple download of pages
		// even if the collector is restarted
		colly.CacheDir("./yandex_cache"),
	)

	// Find and visit all links
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit("https://realty.yandex.ru/")
}
