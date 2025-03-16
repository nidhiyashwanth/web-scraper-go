package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/gocolly/colly"
)

type Product struct {
	Url, Image, Name, Price string
}

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("www.scrapingcourse.com"),
	)
	var products []Product
	var visitedUrls sync.Map

	// Set a custom user agent to avoid being blocked
	// c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36"

	// err := c.SetProxy("https://13.38.153.36:80")
	// if err != nil {
	// 	log.Fatal("Failed to set proxy", err)
	// }

	c.OnHTML("li.product", func(h *colly.HTMLElement) {
		// Scraping logic
		product := Product{}
		product.Url = h.ChildAttr("a", "href")
		product.Image = h.ChildAttr("img", "src")
		product.Name = h.ChildText(".product-name")
		product.Price = h.ChildText(".price")

		products = append(products, product)
	})

	c.OnHTML("a.next", func(h *colly.HTMLElement) {
		nextPage := h.Attr("href")

		if _, found := visitedUrls.Load(nextPage); !found {
			fmt.Println("scraping:", nextPage)
			visitedUrls.Store(nextPage, struct{}{})
			h.Request.Visit(nextPage)
		}
	})

	c.OnScraped(func(r *colly.Response) {
		file, err := os.Create("products.csv")
		if err != nil {
			log.Fatal("Failed to create file", err)
		}

		defer file.Close()

		writer := csv.NewWriter(file)

		// Write the csv headers
		headers := []string{"Url", "Image", "Name", "Price"}

		writer.Write(headers)

		for _, product := range products {
			record := []string{
				product.Url,
				product.Image,
				product.Name,
				product.Price,
			}
			writer.Write(record)
		}
		defer writer.Flush()

	})

	c.Visit("https://www.scrapingcourse.com/ecommerce")

}
