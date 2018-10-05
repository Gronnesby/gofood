package web

import (
	"fmt"
	"strings"

	"github.com/anaskhan96/soup"
)

// Scrape takes the HTML body from a request and prints the contents of the menu ul
func Scrape(body string) map[string][]string {

	doc := soup.HTMLParse(body)
	recipelist := doc.Find("ul", "class", "list-group-flush").FindAll("li", "class", "list-group-item")

	for _, item := range recipelist {
		fmt.Println(strings.TrimSpace(item.Find("h6").Text()))
		for _, recipe := range item.FindAll("li", "class", "recipe") {
			fmt.Println(recipe.Text())
		}
		fmt.Printf("\n\n")
	}

	return nil
}
