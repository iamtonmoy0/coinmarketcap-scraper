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
	writer.Write([]string{"Name", "Symbol", "Market Cap (USD)", "Price (USD)", "Circulating Supply (USD)", "Volume (24h)", "Change (1h)", "Change (24h)", "Change (7d)"})

	// initialize colly
	c := colly.NewCollector(colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/97.0.4692.99 Safari/537.36"))

	c.OnHTML("tbody tr ", func(e *colly.HTMLElement) {
		writer.Write([]string{
			e.ChildText(".cmc-link"),
		})
	})
	c.Visit("https://coinmarketcap.com/all/views/all/")

	log.Printf("Scraping finished, check file %q for results\n", fName)
}
