package main

import (
	"net/http"
	"fmt"
	"github.com/open-sea/coordinator/utils"
)


type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

type sampleFetcher map[string]int

func (f sampleFetcher) Fetch(url string) (string, []string, error) {
	resp, err := http.Get(url)

	// handle the error if there is one
	if err != nil {
		return "", nil, fmt.Errorf("Error fetching")
	}

	// do this now so it won't be forgotten
	defer resp.Body.Close()


	links := utils.GetLinksFromHTML(url, resp.Body)

	//for _, link := range links {
	//	fmt.Println(link)
	//}

	return "", links, nil
}

var m = make(map[string]int)

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:
	if depth <= 0 {
		return
	}

	_, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, u := range urls {
		if _, ok := m[u]; !ok {
			m[u] = 1
			Crawl(u, depth-1, fetcher)
		}
	}
	return
}


var fetcher = sampleFetcher{}

func main() {
	Crawl("https://csh.rit.edu", 5, fetcher)
	for key, _ := range m {
		println(key)
	}
	println(len(m))
}



