package main

import (
	"fmt"
	"net/url"
	"strings"
)

func normalizedURL(inputURL string) (string, error) {
	parsedURL, err := url.Parse(inputURL)
	if err != nil {
		return "", fmt.Errorf("inputURL: %s is not valid", inputURL)
	}
	fullPath := parsedURL.Host + parsedURL.Path

	fullPath = strings.ToLower(fullPath)

	fullPath = strings.TrimSuffix(fullPath, "/")

	return fullPath, nil
}
