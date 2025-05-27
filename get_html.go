package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func getHTML(rawURL string) (string, error) {
	res, err := http.Get(rawURL)
	if err != nil {
		log.Fatal(err)
	}

	contentType := res.Header.Get("content-type")
	if res.StatusCode > 399 {
		return "", fmt.Errorf("status code: %d", res.StatusCode)
	}

	if !strings.Contains(contentType, "text/html") {
		return "", fmt.Errorf("content type not text/html - %s", contentType)
	}

	defer res.Body.Close()
	content, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	return string(content), nil
}
