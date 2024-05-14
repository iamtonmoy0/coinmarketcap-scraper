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
	writer.Write([]string{"Name", "LTP:Last Trade Price", "High", "Low", "CLOSEUP", "YCP:Yesterday Closing Price", "Change", "Trade", "Value", "Volume"})

	// initialize colly
	c := colly.NewCollector()

	c.OnHTML("table tbody", func(e *colly.HTMLElement) {
		writer.Write([]string{
			e.ChildText("tr>td:nth-child(2)>a"),
			e.ChildText("tr>td:nth-child(3)"),
			e.ChildText("tr>td:nth-child(4)"),
			e.ChildText("tr>td:nth-child(5)"),
			e.ChildText("tr>td:nth-child(6)"),
			e.ChildText("tr>td:nth-child(7)"),
			e.ChildText("tr>td:nth-child(8)"),
			e.ChildText("tr>td:nth-child(9)"),
			e.ChildText("tr>td:nth-child(10)"),
			e.ChildText("tr>td:nth-child(11)"),
		})
	})

	// on request
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})
	// on error
	c.OnError(func(e *colly.Response, err error) {
		fmt.Println("something went wrong", err)

	})

	c.Visit("https://www.dsebd.org/latest_share_price_scroll_by_ltp.php")

	log.Printf("Scraping finished, check file %q for results\n", fName)
}
