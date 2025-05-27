package main

import (
	"fmt"
	"net/url"
)

func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int, maxDepth, currentDepth int) {

	if currentDepth > maxDepth {
		fmt.Printf("max depth %d reached, for %s\n", maxDepth, rawCurrentURL)
		return
	}

	structRawBaseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		fmt.Printf("error parsing raw base url %s: %v\n", rawBaseURL, err)
		return
	}

	structCurrentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf("error parsing raw current url %s: %v\n", rawCurrentURL, err)
		return
	}

	if structRawBaseURL.Host != structCurrentURL.Host {
		// fmt.Printf("skipping %s, not in the same domain as %s\n", rawCurrentURL, rawBaseURL)
		return
	}

	// fmt.Printf("DEBUG: Input rawCurrentURL: %s\n", rawCurrentURL)

	absoluteURL, err := staticalizeURL(rawCurrentURL, rawBaseURL)
	if err != nil {
		fmt.Printf("error creating absolute URL %s: %v\n", rawCurrentURL, err)
		return
	}

	// fmt.Printf("DEBUG: After staticalizeURL: %s\n", absoluteURL)

	canonicalURL, err := normalizeURL(absoluteURL)
	if err != nil {
		fmt.Printf("error normalizing URL %s: %v\n", absoluteURL, err)
		return
	}

	// fmt.Printf("DEBUG: After normalizeURL: %s\n", canonicalURL)

	if pages[canonicalURL] >= 1 {
		// fmt.Printf("current URL %s already crawled %d times, skipping\n", absoluteURL, pages[absoluteURL])
		pages[canonicalURL]++

		return
	}

	// fmt.Printf("current URL %s crawling 1st time\n", absoluteURL)
	pages[canonicalURL] = 1

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
		crawlPage(rawBaseURL, u, pages, maxDepth, currentDepth+1)
	}

	fmt.Printf("finished crawling %s\n", absoluteURL)
}
