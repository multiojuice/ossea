package main

import (
	"net/http"
	"fmt"
	"./utils"
)


type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}


// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:
	if depth <= 0 {
		return
	}

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		Crawl(u, depth-1, fetcher)
	}
	return
}


type sampleFetcher map[string]int

func (f sampleFetcher) Fetch(url string) (string, []string, error) {
	fmt.Printf("HTML code of %s ...\n", url)
	resp, err := http.Get(url)
  	
	// handle the error if there is one
	if err != nil {
		return "", nil, fmt.Errorf("Error fetching")
	}

	// do this now so it won't be forgotten
	defer resp.Body.Close()
	
	links := utils.GetLinksFromHTML(url, resp.Body)

	for _, link := range links {
		fmt.Println(link)
	}
	return "", nil, nil
}


var fetcher = sampleFetcher{}

func main() {
	Crawl("https://csh.rit.edu", 4, fetcher)
}



