package main

import (
	"encoding/json"
	"fmt"
	_ "fmt"
	"io/ioutil"
	"strings"

	"github.com/gocolly/colly"
)

type stock_data struct {
	InStock      bool   `json:"inStock"`
	StockMessage string `json:"stockMessage"`
}

// main() contains code adapted from example found in Colly's docs:
// http://go-colly.org/docs/examples/basic/
func main() {
	// Instantiate default collector

	c := colly.NewCollector()

	// On every a element which has href attribute call callback
	//var inStock bool
	stock_check := &stock_data{
		InStock:      false,
		StockMessage: ""}
	//inStock = false

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		// Print link
		if strings.Contains("select delivery location", e.Text) {
			stock_check.InStock = true

		}
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Start scraping on https://hackerspaces.org
	c.Visit("https://www.amazon.com/Microsoft-Xbox-1TB-Compatible-Performance/dp/B09YTP3JS2/ref=sr_1_9?crid=133AZ4XGTLBIQ&keywords=xbox+series+x&qid=1651253940&sprefix=%2Caps%2C460&sr=8-9")

	if stock_check.InStock == true {
		stock_check.StockMessage = "The item is in stock!"
		fmt.Println("The item is in stock!")
	} else if stock_check.InStock == false {
		stock_check.StockMessage = "The item is out of stock!"
		fmt.Println("The item is out of stock!")
	}

	aStringValue, _ := json.Marshal(stock_check)
	fmt.Println(string(aStringValue))

	bytesToWrite := []byte(aStringValue)
	err := ioutil.WriteFile("output.json", bytesToWrite, 0644)
	if err != nil {
		panic(err)
	}
}
