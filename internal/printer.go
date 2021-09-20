package internal

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
	"strings"
)

const baseHabrUrl = "https://habr.com"

type printer struct {
}

type Printer interface {
	Print(data io.Reader) error
}

func (p *printer) isGarbage(title string) bool {
	if title == "<title>" {
		return true
	}
	if title == "Читать далее" ||
		strings.Contains(title, "Читать далее") ||
		strings.Contains(title, "Читать дальше") ||
		strings.Contains(title, "&rarr") {
		return true
	}
	return false
}

func (p *printer) parsePage(r io.Reader) (res map[string]string, err error) {
	res = make(map[string]string, 2)
	z := html.NewTokenizer(r)
	if z == nil {
		return nil, nil
	}
	for {
		tt := z.Next()
		if tt == html.ErrorToken {
			errMsg := z.Err().Error()
			if errMsg == "EOF" {
				return res, nil
			}
			return nil, z.Err()
		}

		if tn, _ := z.TagName(); string(tn) != "a" {
			continue
		}

		isTitleAnchor := false
		var href string
		for {
			k, v, ok := z.TagAttr()
			if !ok {
				break
			}
			attr := string(k)
			if attr == "href" {
				href = string(v)
				if len(href) > 10 && href[0:9] == "/ru/post/" {
					isTitleAnchor = true
				}
				break
			}
		}
		if isTitleAnchor {
			z.Next()
			z.Next()
			title := string(z.Raw())
			if p.isGarbage(title) {
				continue
			}
			res[title] = fmt.Sprintf("%s%s", baseHabrUrl, href)
		}
	}
}

func (p *printer) Print(r io.Reader) error {
	parsed, err := p.parsePage(r)
	if err != nil {
		return err
	}
	for name, link := range parsed {
		fmt.Printf("%s: %s\n", name, link)
	}
	return nil
}

func NewPrinter() Printer {
	return &printer{}
}
