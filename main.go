package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	// fmt.Println(os.Args)
	if len(os.Args) < 2 {
		fmt.Println("no website provided")
		os.Exit(1)
	} else if len(os.Args) > 2 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	url := os.Args[1]
	fmt.Printf("starting crawl of: %s\n", url)

	html, err := getHTML(url)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(html)

	// urls, err := getURLsFromHTML(s, url)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(urls)
}
