package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("no website provided")
		os.Exit(1)
	} else if len(os.Args) > 2 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	url := os.Args[1]
	fmt.Printf("starting crawl of: %s\n", url)

	maxDepth := 10
	pages := make(map[string]int)
	crawlPage(url, url, pages, maxDepth, 0)

	for page, depth := range pages {
		fmt.Printf("%d - %s\n", depth, page)
	}

}
