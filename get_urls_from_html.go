package main

import (
	"fmt"
	"net/url"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func staticalizeURL(urlToStaticalize, baseURL string) (string, error) {
	structURL, err := url.Parse(urlToStaticalize)
	if err != nil {
		return "", err
	}

	if structURL.Scheme == "" || structURL.Host == "" {
		baseStructURL, err := url.Parse(baseURL)
		if err != nil {
			return "", err
		}
		toReturn := fmt.Sprintf("%s://%s%s", baseStructURL.Scheme, baseStructURL.Host, structURL.Path)
		return toReturn, nil
	}
	return urlToStaticalize, nil
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
						return nil, err
					}
					slice = append(slice, s)
				}
			}
		}
	}
	return slice, nil
}
