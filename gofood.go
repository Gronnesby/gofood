package main

import (
	"flag"

	"github.com/gronnesby/gofood/web"
)

var loc = flag.String("loc", "", "Set the location to get the menu from, defaults to Graduate College")

func main() {
	flag.Parse()

	body := web.GetWebpage()
	web.Scrape(body)
}
