package main

import (
	"encoding/csv"
	"log"
	"os"

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

	c.OnHTML("li.product", func(h *colly.HTMLElement) {
		// Scraping logic
		product := Product{}
		product.Url = h.ChildAttr("a", "href")
		product.Image = h.ChildAttr("img", "src")
		product.Name = h.ChildText(".product-name")
		product.Price = h.ChildText(".price")

		products = append(products, product)
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
