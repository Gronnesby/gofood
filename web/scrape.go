package web

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/anaskhan96/soup"
	"github.com/fatih/color"
)

type Menu struct {
	MealType string
	Opt      []string
	NumItems int
}

// Scrape takes the HTML body from a request and prints the contents of the menu ul
func Scrape(body string, unlimstr bool, tabprint bool) map[string][]string {

	w := new(tabwriter.Writer)

	// Format in tab-separated columns with a tab stop of 8.
	w.Init(os.Stdout, 0, 8, 0, '\t', 0)

	doc := soup.HTMLParse(body)
	menu := make(map[string][]string)
	location := strings.TrimSpace(doc.Find("div", "id", "jumbotronFeature").Find("h1").Text())

	// Print the location in bold yellow text
	color.Set(color.FgGreen, color.Bold)
	fmt.Printf("<<<- %s ->>>\n\n", location)
	color.Unset()

	// Find all the meal cards
	meals := doc.FindAll("div", "class", "mealCard")

	// Go through all the meals of the day
	for _, m := range meals {
		if m.Error != nil {
			fmt.Printf("No Data Available\n")
			return nil
		}

		// Print the meal type in bold
		color.Set(color.FgHiYellow, color.Bold)
		fmt.Printf("--- %s ---\n", m.Find("div", "class", "card-header").Find("h5", "class", "mealName").Text())
		color.Unset()

		var menus []Menu
		var rows int

		// Iterate through all the courses for the meal
		recipes := m.Find("ul", "class", "list-group")
		for i, item := range recipes.FindAll("li", "class", "list-group-item") {
			if item.Error != nil {
				panic(item.Error)
			}
			category := strings.TrimSpace(item.Find("h6").Text())
			options := item.FindAll("li", "class", "recipe")

			menus = append(menus, Menu{
				MealType: category,
				NumItems: len(options),
			})

			if menus[i].NumItems > rows {
				rows = len(options)
			}

			for _, recipe := range options {
				// if unlimstr || len(recipe.Text()) < 20 {
				// 	menus[i].Opt = append(menus[i].Opt, recipe.Text())
				// } else {
				// 	menus[i].Opt = append(menus[i].Opt, fmt.Sprintf("%s%s", recipe.Text()[0:20], "... etc"))
				// }
				menus[i].Opt = append(menus[i].Opt, recipe.Text())

			}
		}

		if tabprint {
			Tabprint(menus, rows)
		} else {
			Seqprint(menus)
		}
	}

	fmt.Printf("\n\n")
	return menu
}
