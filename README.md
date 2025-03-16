# Web Scraper Go

A simple and efficient web scraper built with Go that extracts product information from an e-commerce website and saves it to a CSV file.

## Author

**Nidhi Yashwanth**

- GitHub: [github.com/nidhiyashwanth](https://github.com/nidhiyashwanth)

## Features

- Scrapes product details including:
  - Product URL
  - Product Image URL
  - Product Name
  - Product Price
- Handles pagination automatically
- Prevents duplicate URL visits
- Exports data to CSV format
- Concurrent scraping using Colly framework

## Prerequisites

- Go 1.16 or higher
- Git

## Dependencies

The project uses the following external package:

- [Colly](https://github.com/gocolly/colly) - Elegant Scraping Framework for Golang

## Installation

1. Clone the repository:

```bash
git clone [your-repository-url]
cd web-scraper-go
```

2. Install dependencies:

```bash
go mod init web-scraper-go
go get github.com/gocolly/colly
```

## Usage

1. Run the scraper:

```bash
go run scraper.go
```

The script will:

- Start scraping from the initial URL
- Process all available product pages
- Save the results in `products.csv` in the same directory

## Output

The scraper generates a CSV file (`products.csv`) with the following columns:

- Url: The product's page URL
- Image: URL of the product image
- Name: Product name
- Price: Product price

## Configuration

The scraper comes with some configurable options (currently commented out in the code):

- Custom User Agent
- Proxy support

To enable these features, uncomment the relevant sections in the code.

## Error Handling

The scraper includes basic error handling for:

- File operations
- Network requests
- Data parsing

## Contributing

Feel free to submit issues and enhancement requests!
