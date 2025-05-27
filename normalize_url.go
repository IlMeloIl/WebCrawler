package main

import (
	"net/url"
	"strings"
)

func normalizeURL(rawURL string) (string, error) {
	urlStruct, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}

	normalizedURL := urlStruct.Host + urlStruct.Path
	normalizedURL = strings.ToLower(normalizedURL)
	normalizedURL = strings.TrimSuffix(normalizedURL, "/")

	return normalizedURL, nil
}
