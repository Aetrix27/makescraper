package main

import (
	"fmt"
	_ "fmt"
	"strings"

	"github.com/gocolly/colly"
)

type data struct {
	json string
}

// main() contains code adapted from example found in Colly's docs:
// http://go-colly.org/docs/examples/basic/
func main() {
	// Instantiate default collector

	c := colly.NewCollector()

	// On every a element which has href attribute call callback
	var inStock bool
	inStock = false

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {

		//link := e.Attr("href")

		// Print link
		//fmt.Printf("Link found: %q -> %s\n", e.Text, link)
		//_ = inStock

		if strings.Contains("select delivery location", e.Text) {
			inStock = true

		}
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Start scraping on https://hackerspaces.org
	c.Visit("https://www.amazon.com/Microsoft-Xbox-1TB-Compatible-Performance/dp/B09YTP3JS2/ref=sr_1_9?crid=133AZ4XGTLBIQ&keywords=xbox+series+x&qid=1651253940&sprefix=%2Caps%2C460&sr=8-9")

	if inStock == true {
		fmt.Println("The item is in stock!")
	} else if inStock == false {
		fmt.Println("The item is out of stock!")
	}

}
