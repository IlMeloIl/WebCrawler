package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"strconv"
	"sync"
)

type config struct {
	pages              map[string]int
	baseURL            *url.URL
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
}

func newConfig(baseURLString string, maxConcurrency int) (*config, error) {

	parsedBaseURL, err := url.Parse(baseURLString)
	if err != nil {
		return nil, fmt.Errorf("error parsing base URL %s: %v", baseURLString, err)
	}

	return &config{
		pages:              make(map[string]int),
		baseURL:            parsedBaseURL,
		mu:                 &sync.Mutex{},
		concurrencyControl: make(chan struct{}, maxConcurrency),
		wg:                 &sync.WaitGroup{},
	}, nil
}

func main() {

	if len(os.Args) < 4 {
		fmt.Println("provide website, max concurrency, and max depth to crawl")
		os.Exit(1)
	} else if len(os.Args) > 4 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	url := os.Args[1]

	maxConcurrency, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal("error parsing max concurrency to int: ", err)
	}
	if maxConcurrency <= 0 {
		log.Fatal("max concurrency must be greater than 0")
	}

	maxDepth, err := strconv.Atoi(os.Args[3])
	if err != nil {
		log.Fatal("error parsing max depth to int: ", err)
	}
	if maxDepth <= 0 {
		log.Fatal("max depth must be greater than 0")
	}

	cfg, err := newConfig(url, maxConcurrency)
	if err != nil {
		log.Fatal("error creating new config: ", err)
	}

	fmt.Printf("starting crawl of: %s\n", url)

	cfg.wg.Add(1)
	go func(startURL string) {
		defer cfg.wg.Done()
		cfg.concurrencyControl <- struct{}{}
		defer func() { <-cfg.concurrencyControl }()

		cfg.crawlPage(startURL, maxDepth, 0)
	}(url)

	cfg.wg.Wait()

	cfg.mu.Lock()

	for page, depth := range cfg.pages {
		fmt.Printf("%d - %s\n", depth, page)
	}
	cfg.mu.Unlock()
}
