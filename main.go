package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector()

	c.OnHTML(".next.page-numbers", func(e *colly.HTMLElement) {
		href := e.Attr("href")
		fmt.Println("--------------href:", href)
		e.Request.Visit(href)

	})

	c.OnHTML(".entry-title", func(e *colly.HTMLElement) {
		title := e.Text
		fmt.Println("title:", title)
	})

	c.OnRequest(func(r *colly.Request) {
		URL := r.URL
		fmt.Println("URL", URL)
	})

	c.Visit("http://nghiahsgs.com")

}
