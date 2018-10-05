package web

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const (
	foodurl = "https://menus.princeton.edu/dining/_Foodpro/online-menu/menuDetails.asp?myaction=read&sName=Princeton+University+Campus+Dining&dtdate=10%2F3%2F2018&locationNum=04&locationName=%20+Graduate+College+&naFlag=1"
)

// GetWebpage sets todays date in the URL query and performs a Get request.
// Function returns the HTML body from the Get request.
func GetWebpage() string {

	u, err := url.Parse(foodurl)
	if err != nil {
		panic(err)
	}
	q := u.Query()
	currentTime := time.Now()
	date := fmt.Sprintf("%d/%d/%d", currentTime.Month(), currentTime.Day(), currentTime.Year())
	q.Set("dtdate", date)
	u.RawQuery = q.Encode()

	resp, err := http.Get(u.String())
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return string(body)
}
