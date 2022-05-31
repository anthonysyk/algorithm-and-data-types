package golang

import (
	"fmt"
	"sync"
	"testing"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

type History struct {
	history map[string]bool
	mu      sync.Mutex
}

func NewHistory() History {
	return History{history: make(map[string]bool)}
}

func (c *History) isFetched(url string) bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	if _, ok := c.history[url]; !ok {
		c.history[url] = true
		return false
	}

	return true
}

func Crawl(history *History, url string, depth int, fetcher Fetcher, wg *sync.WaitGroup) {
	defer wg.Done()
	isFetched := history.isFetched(url)
	if depth <= 0 || isFetched {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println("error", err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		wg.Add(1)
		go Crawl(history, u, depth-1, fetcher, wg)
	}
}

func syncCrawl(url string, depth int, fetcher Fetcher) {
	history := NewHistory()
	var wg sync.WaitGroup
	wg.Add(1)
	go Crawl(&history, url, depth, fetcher, &wg)
	wg.Wait()
}

func TestCrawl(t *testing.T) {
	syncCrawl("https://golang.org/", 4, fetcher)
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
