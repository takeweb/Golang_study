package main

import (
	"github.com/PuerkitoBio/goquery"
)

func main() {
	p := "https://golang.org"
	doc, er := goquery.NewDocument(p)
	if er != nil {
		panic(er)
	}

	doc.Find("a").Each(func(n int, sel *goquery.Selection) {
		lk, _ := sel.Attr("href")
		println(n, sel.Text(), "(", lk, ")")
	})
}
