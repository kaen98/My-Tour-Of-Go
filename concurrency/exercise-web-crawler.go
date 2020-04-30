package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

type CrawlUrls struct {
	urlsMap map[string]int
	mux sync.Mutex
	wg sync.WaitGroup
}

func (c CrawlUrls) checkUrlIsExists(url string) bool {
	c.mux.Lock()
	defer c.mux.Unlock()
	if _,ok := c.urlsMap[url]; !ok {
		c.urlsMap[url] = 1
		return false
	}else{
		return true
	}
}

var crawlUrls CrawlUrls = CrawlUrls{urlsMap: make(map[string]int)}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	defer crawlUrls.wg.Done()

	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	if crawlUrls.checkUrlIsExists(url) == true {
		return
	}

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
		crawlUrls.wg.Add(1)
		go Crawl(u, depth-1, fetcher)
	}
	return
}

func main() {
	crawlUrls.wg.Add(1)
	Crawl("https://golang.org/", 4, fetcher)
	crawlUrls.wg.Wait()
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
