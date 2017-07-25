package main

import (
	"github.com/yhat/scrape"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

type TraderDetails struct {
	Url string

	Name    string
	Summary string
	Images  []string
}

func NewTraderDetails(url string) (t TraderDetails) {
	return TraderDetails{
		Url: url,
	}
}

func matchTraderDescription(n *html.Node) bool {
	if n.DataAtom == atom.Div && scrape.Attr(n, "class") == "content" {
		return scrape.Attr(n, "class") == "content"
	}

	return false
}

func getTraderDetails(url string, chTraderDetails chan TraderDetails) {
	t := NewTraderDetails(url)

	root, err := fetchAndParse(url)
	if err != nil {
		panic(err)
	}

	slider, ok := scrape.Find(root, scrape.ByClass("slider"))
	if ok {
		imgs := scrape.FindAll(slider, scrape.ByClass("slider__img"))
		for _, img := range imgs {
			imgUrl := scrape.Attr(img, "zrs-src")
			t.Images = append(t.Images, imgUrl)
		}

		if traderName, ok := scrape.Find(slider, scrape.ByClass("slider__title")); ok {
			t.Name = scrape.Text(traderName)
		}
	}

	traderSummary, ok := scrape.Find(root, scrape.ByClass("content"))
	if ok {
		t.Summary = scrape.Text(traderSummary)
	}

	chTraderDetails <- t

	return
}
