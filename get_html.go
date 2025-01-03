package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func GetHMTL(rawURL string) (string, error) {
	res, err := http.Get(rawURL)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	if res.StatusCode > 499 {
		return "", fmt.Errorf("got Server error: %s", res.Status)
	}

	if res.StatusCode > 399 {
		return "", fmt.Errorf("got HTTP error: %s", res.Status)
	}

	contentType := res.Header.Get("content-type")
	if !strings.Contains(contentType, "text/html") {
		return "", fmt.Errorf("got non-HTML response: %s in %s", contentType, rawURL)
	}

	htmlBodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("couldn't read body: %v", err)
	}

	return string(htmlBodyBytes), nil
}
