package internal

import (
	"fmt"
	"net/http"
)

const baseHabrPageUrl = "https://habr.com/ru/all/"

type scrapper struct {
	printer Printer
}

type Scrapper interface {
	Scrap(top int, pages int) error
}

func (s *scrapper) buildHabrUrl(top int) string {
	return fmt.Sprintf("%stop%d/", baseHabrPageUrl, top)
}

func (s *scrapper) retrieveData(url string, pages int) error {
	for i := 1; i < pages; i++ {
		currUrl := url
		if i != 1 {
			currUrl = fmt.Sprintf("%spage%d/", url, i)
		}
		resp, err := http.Get(currUrl)
		if err != nil {
			return err
		}
		if err := s.printer.Print(resp.Body); err != nil {
			return err
		}
	}
	return nil
}

func (s *scrapper) Scrap(top int, pages int) error {
	url := s.buildHabrUrl(top)
	return s.retrieveData(url, pages)
}

func NewScrapper(printer Printer) Scrapper {
	return &scrapper{printer}
}
