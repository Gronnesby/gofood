package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/fatih/color"

	"github.com/gronnesby/gofood/web"
)

var locations = []web.LocationParams{
	{"01", "Rockefeller & Mathey Colleges"},
	{"02", "Butler & Wilson Colleges"},
	{"03", "Forbes College"},
	{"04", "Graduate College"},
	{"05", "Center for Jewish Life"},
	{"08", "Whitman College"},
}

var loc = flag.String("loc", "", "Specify a dining location.")
var unlimstr = flag.Bool("unlimstr", false, "Print unlimited string length in the cells, may offset table.")
var tomorrow = flag.Bool("t", false, "Get the menu for tomorrow instead of today.")
var tabprint = flag.Bool("tab", false, "Prints the results in a table. Might soft wrap for long menus.")

func main() {
	flag.Parse()

	var query []web.LocationParams

	if *loc != "" {
		if *loc == "all" {
			color.Set(color.FgMagenta, color.Bold)
			fmt.Println("Available Locations")
			color.Unset()
			for _, l := range locations {
				fmt.Println(l.LocationName)
			}
			return
		}

		for _, l := range locations {
			if strings.Contains(strings.ToLower(l.LocationName), *loc) {
				query = append(query, l)
			}
		}
		if len(query) == 0 {
			log.Fatal(fmt.Sprintf("Location %s not available", *loc))
		}
	} else {
		query = locations
	}

	responses := web.GetWebpage(query, *tomorrow)
	if *tomorrow {
		color.Set(color.Bold)
		fmt.Printf("\nTomorrows Menu\n\n")
		color.Unset()
	}
	for content := range responses {
		web.Scrape(content, *unlimstr, *tabprint)
	}
}
