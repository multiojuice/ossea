package utils

import (
	"golang.org/x/net/html"
	"io"
	"regexp"
	s "strings"
)

func GetLinksFromHTML(url string, body io.Reader) ([]string) {
	var links []string
	z := html.NewTokenizer(body)
	var getDomain = regexp.MustCompile(`^(?:https?:\/\/)?(?:[^@\/\n]+@)?(?:www\.)?([^:\/?\n]+)`)
	domains := getDomain.FindStringSubmatch(url)
	if len(domains) < 1 {
		return []string{}
	}
	var baseUrl = domains[0]

	for {
		tt := z.Next()

		switch tt {
		case html.ErrorToken:
			cleanedLinks := removeUnknownDomains(url, links);
			return cleanedLinks
		case html.StartTagToken, html.EndTagToken:
			token := z.Token()
			if token.Data == "a" {
				for _, attr := range token.Attr {
					if attr.Key == "href" {
						formattedLink, err := formatLink(baseUrl, url, attr.Val)
						if !err {
							links = append(links, formattedLink)
						}
					}
				}
			}
		}
	}
}


func removeUnknownDomains(url string, links []string) ([]string) {
	var knownDomain = regexp.MustCompile(`(http:\/\/|www\.|https:\/\/)([[:ascii:]]*\.|)csh.rit.edu(\/[[:ascii:]]*|)`)
	var cleanLinks []string
	for _, s := range links {
		if knownDomain.MatchString(s) {
			cleanLinks = append(cleanLinks, s)
		}
	}
	return cleanLinks
}


func formatLink(baseUrl string, currentUrl string, link string) (string, bool) {
	if s.HasPrefix(link, "https://") || s.HasPrefix(link, "http://") {
		return link, false
	} else if s.HasPrefix(link, "./") {
		return currentUrl + link[1:], false
	} else if s.HasPrefix(link, "/") {
		return baseUrl + link, false
	} else {
		return "", true
	}
}
