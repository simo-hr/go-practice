package main

import (
	"fmt"
	"sync"
)

// Fetcher インターフェース
type Fetcher interface {
	// Fetch は URL の本文とそのページで見つかった URL のスライスを返します。
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl は fetcher を使って、url から始まり、最大 depth まで再帰的にページをクロールします。
func Crawl(url string, depth int, fetcher Fetcher, wg *sync.WaitGroup, cache *sync.Map) {
	defer wg.Done()

	if depth <= 0 {
		return
	}

	// 同じURLを2度取得しないようにする
	if _, loaded := cache.LoadOrStore(url, true); loaded {
		return
	}

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)

	// 子URLを並列にクロール
	for _, u := range urls {
		wg.Add(1)
		go Crawl(u, depth-1, fetcher, wg, cache)
	}
}

func main() {
	var wg sync.WaitGroup
	cache := &sync.Map{}

	wg.Add(1)
	Crawl("https://golang.org/", 4, fetcher, &wg, cache)
	wg.Wait()
}

// fakeFetcher はフェッチしたURLに対応する事前に用意された結果を返す Fetcher です。
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

// fetcher は事前に用意された fakeFetcher です。
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
