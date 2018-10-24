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

type SafeMap struct {
	m       map[string]bool
	urlBody map[string]string
	mux     sync.RWMutex
}

func (sm *SafeMap) ToFetch(url string) {
	sm.mux.Lock()
	sm.m[url] = false
	sm.mux.Unlock()
}

func (sm *SafeMap) Fetched(url string, body string) {
	sm.mux.Lock()
	sm.m[url] = true
	sm.urlBody[url] = body
	sm.mux.Unlock()
}

func (sm *SafeMap) IsFetched(url string) bool {
	sm.mux.Lock()
	defer sm.mux.Unlock()
	return sm.m[url]
}

func (sm *SafeMap) IsEntirelyFetched() bool {
	sm.mux.Lock()
	defer sm.mux.Unlock()
	for url, _ := range sm.m {
		if !sm.m[url] {
			return false
		}
	}
	return true
}

func (sm *SafeMap) NextToFetch() string {
	sm.mux.Lock()
	defer sm.mux.Unlock()
	for url, _ := range sm.m {
		if !sm.m[url] {
			return url
		}
	}
	return ""
}

func (sm *SafeMap) String() {
	sm.mux.Lock()
	defer sm.mux.Unlock()
	for url, body := range sm.urlBody {
		fmt.Printf("found: %s : %s\n", url, body)
	}
}

var sm = SafeMap{m: make(map[string]bool), urlBody: make(map[string]string)}
var wg sync.WaitGroup

func Crawl(url string, depth int, fetcher Fetcher) {
	//Crawling an URL for a specified depth
	//urlToFetch := make(chan string)
	var insideCrawl func(url string, fetcher Fetcher, depth int, sm SafeMap)
	insideCrawl = func(url string, fetcher Fetcher, depth int, sm SafeMap) {
		defer wg.Done()

		if depth <= 0 {
			fmt.Println("Depth 0 reached for url ", url)
			return
		}
		//body, urls, err := fetcher.Fetch(url)
		if sm.IsFetched(url) {
			return
		}
		body, urls, err := fetcher.Fetch(url)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("found: %s %q %q\n", url, body, urls)
		sm.Fetched(url, body)
		for _, u := range urls {
			wg.Add(1)
			go insideCrawl(u, fetcher, depth-1, sm)
		}
	}
	wg.Add(1)
	insideCrawl(url, fetcher, depth, sm)
	wg.Wait()
	return
}

func main() {
	Crawl("https://golang.org/", 4, fetcher)
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
