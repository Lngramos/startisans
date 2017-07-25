package main

import (
	"log"
	"net/http"

	"golang.org/x/net/html"
)

func fetchAndParse(url string) (root *html.Node, err error) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	b := resp.Body
	defer b.Close()

	root, err = html.Parse(b)
	if err != nil {
		panic(err)
	}

	return
}

func main() {
	urls, err := getTraderUrls("http://www.startisans.net/whats-on/")
	if err != nil {
		panic(err)
	}

	chTraderDetails := make(chan TraderDetails)

	for _, url := range urls {
		go getTraderDetails(url, chTraderDetails)
	}

	for c := 0; c < len(urls); {
		select {
		case details := <-chTraderDetails:
			log.Println(details)
			c++
		}
	}

	close(chTraderDetails)
}
