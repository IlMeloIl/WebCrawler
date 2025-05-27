package main

import (
	"log"
	"net/url"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func staticalizeURL(href, base string) (string, error) {
	hrefURL, err := url.Parse(href)
	if err != nil {
		return "", err
	}

	if hrefURL.IsAbs() {
		return href, nil
	}

	baseURL, err := url.Parse(base)
	if err != nil {
		return "", err
	}

	absoluteURL := baseURL.ResolveReference(hrefURL)

	return absoluteURL.String(), nil
}

func getURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {
	doc, err := html.Parse(strings.NewReader(htmlBody))
	if err != nil {
		return nil, err
	}

	var slice []string

	for n := range doc.Descendants() {
		if n.Type == html.ElementNode && n.DataAtom == atom.A {
			for _, a := range n.Attr {
				if a.Key == "href" {
					s, err := staticalizeURL(a.Val, rawBaseURL)
					if err != nil {
						log.Printf("error staticalizing URL %s: %v\n", a.Val, err)
						continue
					}
					slice = append(slice, s)
				}
			}
		}
	}
	return slice, nil
}
