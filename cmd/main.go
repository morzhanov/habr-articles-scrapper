package main

import (
	"flag"
	"github.com/morzhanov/habr-articles-scrapper/internal"
	"log"
)

var top, pages int

func init() {
	const (
		defaultTop   = 10
		defaultPages = 20
	)
	flag.IntVar(&top, "top", defaultTop, "top results")
	flag.IntVar(&pages, "pages", defaultPages, "count of pages")
}

func main() {
	flag.Parse()
	printer := internal.NewPrinter()
	scrapper := internal.NewScrapper(printer)
	if err := scrapper.Scrap(top, pages); err != nil {
		log.Fatal(err)
	}
}
