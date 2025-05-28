package main

import (
	"fmt"
	"net/url"
)

func (cfg *config) addPageVisit(normalizedURL string) (isFirstVisit bool) {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()

	_, exists := cfg.pages[normalizedURL]
	if exists {
		cfg.pages[normalizedURL]++
		return false
	}

	cfg.pages[normalizedURL] = 1
	return true
}

func (cfg *config) crawlPage(rawCurrentURL string, maxDepth, currentDepth int) {

	if currentDepth > maxDepth {
		fmt.Printf("max depth %d reached, for %s\n", maxDepth, rawCurrentURL)
		return
	}

	structCurrentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf("error parsing raw current url %s: %v\n", rawCurrentURL, err)
		return
	}

	if cfg.baseURL.Host != structCurrentURL.Host {
		return
	}

	absoluteURL, err := staticalizeURL(rawCurrentURL, cfg.baseURL.String())
	if err != nil {
		fmt.Printf("error creating absolute URL %s: %v\n", rawCurrentURL, err)
		return
	}

	canonicalURL, err := normalizeURL(absoluteURL)
	if err != nil {
		fmt.Printf("error normalizing URL %s: %v\n", absoluteURL, err)
		return
	}

	if !cfg.addPageVisit(canonicalURL) {
		return
	}

	html, err := getHTML(absoluteURL)
	if err != nil {
		fmt.Printf("error getting HTML for %s: %v\n", absoluteURL, err)
		return
	}

	urls, err := getURLsFromHTML(html, absoluteURL)
	if err != nil {
		fmt.Printf("error getting URLs from HTML for %s: %v\n", absoluteURL, err)
	}

	for _, u := range urls {
		if currentDepth+1 <= maxDepth {
			cfg.wg.Add(1)

			nextURL := u
			go func(urlToCrawl string, newDepth int) {

				defer cfg.wg.Done()

				cfg.concurrencyControl <- struct{}{}

				defer func() { <-cfg.concurrencyControl }()

				cfg.crawlPage(urlToCrawl, maxDepth, newDepth)
			}(nextURL, currentDepth+1)
		}
	}

	fmt.Printf("finished crawling %s\n", absoluteURL)
}
