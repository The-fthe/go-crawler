package main

import (
	"fmt"
	"net/url"
)

func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) {
	parseCurrURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Println("Parse Curr Url error: ", err.Error())
		return
	}
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		fmt.Println("Parse Base Url error: ", err.Error())
		return
	}

	if baseURL.Hostname() != parseCurrURL.Hostname() {
		return
	}

	normalizeURL, err := normalizedURL(rawCurrentURL)
	if err != nil {
		fmt.Println("Normalize error: ", err.Error())
		return
	}

	if _, visited := pages[normalizeURL]; visited {
		pages[normalizeURL]++
		return
	}

	pages[normalizeURL] = 1

	fmt.Println("Crawl url: ", rawCurrentURL)

	htmlBody, err := GetHMTL(rawCurrentURL)
	if err != nil {
		fmt.Println("GetHTML error: ", err.Error())
		return
	}
	fmt.Println("htmlBody: ", htmlBody[:20])

	nextUrls, err := getURLsFromHTML(htmlBody, rawBaseURL)
	if err != nil {
		fmt.Println("GetURLs error: ", err.Error())
		return
	}

	for _, url := range nextUrls {
		crawlPage(rawBaseURL, url, pages)
	}
}
