package web

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sync"
	"time"
)

// LocationParams represents the URL parameters for a dining location
type LocationParams struct {
	LocationNum  string
	LocationName string
}

const (
	foodurl = "https://menus.princeton.edu/dining/_Foodpro/online-menu/menuDetails.asp?myaction=read&sName=Princeton+University+Campus+Dining&naFlag=1"
)

// GetWebpage sets todays date in the URL query and performs a Get request.
// Function returns the HTML body from the Get request.
func GetWebpage(locations []LocationParams, tomorrow bool) chan string {

	u, err := url.Parse(foodurl)
	if err != nil {
		panic(err)
	}
	currentTime := time.Now()
	if tomorrow {
		currentTime = currentTime.AddDate(0, 0, 1)
	}
	date := fmt.Sprintf("%d/%d/%d", currentTime.Month(), currentTime.Day(), currentTime.Year())

	content := make(chan string)
	var wg sync.WaitGroup
	for i, loc := range locations {
		wg.Add(1)
		go func(loc LocationParams, i int) {
			defer wg.Done()
			q := u.Query()
			q.Set("dtdate", date)
			q.Set("locationName", loc.LocationName)
			q.Set("locationNum", loc.LocationNum)
			u.RawQuery = q.Encode()

			response, err := http.Get(u.String())
			if err != nil {
				panic(err)
			}
			defer response.Body.Close()
			body, err := ioutil.ReadAll(response.Body)
			if err != nil {
				panic(err)
			}
			content <- string(body)
		}(loc, i)
	}

	go func() {
		wg.Wait()
		close(content)
	}()

	return content
}
