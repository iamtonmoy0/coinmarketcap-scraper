package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
)

func main() {
	fName := "data.csv"
	file, err := os.Create(fName)
	if err != nil {
		fmt.Println("failed to write file")
		return
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()
	// write csv header
	writer.Write([]string{"Name"})

	// initialize colly
	c := colly.NewCollector()

	c.OnHTML("tbody  ", func(e *colly.HTMLElement) {
		result := e.ChildText(".cmc-table__column-name--name")
		if result == "" {
			fmt.Println("No result found for .cmc-table__column-name--name")
		} else {
			fmt.Println(result)
			writer.Write([]string{result})
		}
	})
	// on request
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})
	// on error
	c.OnError(func(e *colly.Response, err error) {
		fmt.Println("something went wrong", err)
	})

	c.Visit("https://coinmarketcap.com/all/views/all/")

	log.Printf("Scraping finished, check file %q for results\n", fName)
}
