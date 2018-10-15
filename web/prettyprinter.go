package web

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/fatih/color"
)

func Tabprint(menus []Menu, rows int) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 4, ' ', tabwriter.Debug)

	for _, v := range menus {
		fmt.Fprintf(w, "%s\t", v.MealType)
	}
	fmt.Fprintf(w, "\n")

	for i := 0; i < rows; i++ {
		for _, v := range menus {
			if len(v.Opt) > i {
				fmt.Fprintf(w, "%s\t", v.Opt[i])
			} else {
				fmt.Fprintf(w, "\t")
			}
		}
		fmt.Fprintf(w, "\n")
	}
	w.Flush()
	fmt.Printf("\n\n")
}

func Seqprint(menus []Menu) {

	for _, v := range menus {
		color.Set(color.FgMagenta, color.Bold)
		fmt.Printf("%s\t", v.MealType)
		color.Unset()

		fmt.Printf("\n")
		for _, r := range v.Opt {
			fmt.Printf("%s\n", r)
		}
		fmt.Printf("\n")
	}

}
