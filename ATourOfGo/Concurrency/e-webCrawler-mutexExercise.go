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

type fakeResult struct {
	body string
	urls []string
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

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

type SafeCrawlCounter struct {
	mu sync.Mutex
	v  map[string]bool
}

func (c *SafeCrawlCounter) SetValue(key string) {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	c.v[key] = true
	c.mu.Unlock()
}

func (c *SafeCrawlCounter) GetValue(key string) bool {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	defer c.mu.Unlock() // since only after the process has received the value we must unlock it, we use defer to achieve this
	return c.v[key]
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, c *SafeCrawlCounter) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice. - Done
	// This implementation doesn't do either:
	if depth <= 0 {
		return
	}
	visited := c.GetValue(url)
	// fmt.Printf("Visited %v, %v\n", url, visited)
	if visited == false {
		c.SetValue(url)
		body, urls, err := fetcher.Fetch(url)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("found: %s %q\n", url, body)
		for _, u := range urls {
			// fmt.Println("Next : ", u)
			Crawl(u, depth-1, fetcher, c)
		}
	} else {
		return
	}
}

func main() {
	var c SafeCrawlCounter
	c = SafeCrawlCounter{v: make(map[string]bool)}
	Crawl("https://golang.org/", 4, fetcher, &c)
}
