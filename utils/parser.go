package utils

import (
	"io"
	 s "strings"
	"golang.org/x/net/html"
)

func GetLinksFromHTML(url string, body io.Reader) ([]string) {
	var links []string
	z := html.NewTokenizer(body)

	for {
		tt := z.Next()

		switch tt {
		case html.ErrorToken:
			return links
		case html.StartTagToken, html.EndTagToken:
			token := z.Token()
			if token.Data == "a" {
				for _, attr := range token.Attr {
					if attr.Key == "href" {
						formattedLink := formatLink(url, attr.Val)
						links = append(links, formattedLink)
						
					}
				}
			}
		}
	}

	return links
}


func formatLink(baseUrl string, link string) (string) {
	if s.HasPrefix(link, "https://") || s.HasPrefix(link, "http://") {
		return link
	} else if s.HasPrefix(link, "./") {
		return baseUrl + link[1:]
	} else if s.HasPrefix(link, "/") {
		return baseUrl + link
	} else {
		return baseUrl + "/" + link
	}
}
