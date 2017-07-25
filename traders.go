package main

import (
	"github.com/yhat/scrape"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func traderUrlsMatcher(n *html.Node) bool {
	if n.DataAtom == atom.A && n.Parent.DataAtom == atom.Div {
		return scrape.Attr(n.Parent, "class") == "post post--traders"
	}

	return false
}

func getTraderUrls(url string) (urls []string, err error) {
	root, err := fetchAndParse(url)
	if err != nil {
		panic(err)
	}

	traders := scrape.FindAll(root, traderUrlsMatcher)
	for _, article := range traders {
		urls = append(urls, scrape.Attr(article, "href"))
	}

	return
}
