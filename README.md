# nordnet-api-client

Go-library for interacting with Nordnet API. Nordnet does not offer publicly available API's in all nordic countries so
this library accesses the API as a regular web browser would. (By using cookies and other special HTTP headers).
It then converts JSON responses to go-structs for programmatic use.

Currently its only possible to retrieve user's stock watchlists, but the development has begun! :-)

## Usage example

1. Open your browser's developer console -> Go to "Network" (it might be called something else)
2. Go to https://nordnet.fi -> Login
3. From the http response copy the values of "Set-Cookie" and "ntag" headers (you can also copy cookies and ntag from any further requests that browser has sent)
4. Write your example program

```go

package main

import (
	"fmt"

	api "github.com/phasi/nordnet-api-client"
)

var cookie string = "<paste your cookies here, remember to escape any possible quotation marks>"
var nTag string = "<copy the ntag value here, that is used to identify you together with the cookies>"

func main() {
	// This example prints user's stock watch list names and IDs.
	api.DefaultDriver.MandatoryHeaders.Add("Cookie", cookie)
	api.DefaultDriver.MandatoryHeaders.Add("ntag", nTag)
	client := api.NewNordnetClient(api.DefaultDriver)

	watchListDetail, err := client.GetWatchLists()
	if err != nil {
		panic(err)
	}
	fmt.Println(watchListDetail)
}

```

Happy hacking!