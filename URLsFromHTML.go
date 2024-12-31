package main

import (
	"golang.org/x/net/html"
	"net/url"
	"strings"
)

func getURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {
	URLs := []string{}
	normalizeURLs := []string{}

	htmlReader := strings.NewReader(htmlBody)
	node, err := html.Parse(htmlReader)
	if err != nil {
		return URLs, err
	}
	findLink(node, &URLs)
	for _, rawURL := range URLs {
		parsedURL, err := url.Parse(rawURL)
		if err != nil {
			return nil, err
		}
		absoleteURL := ""
		if parsedURL.Host == "" {
			absoleteURL = rawBaseURL + parsedURL.Path
		} else {
			absoleteURL = parsedURL.String()
		}
		normalizeURLs = append(normalizeURLs, absoleteURL)
	}

	return normalizeURLs, nil
}

func findLink(node *html.Node, urlLinks *[]string) {
	if node.Type == html.ElementNode && node.Data == "a" {
		for _, attr := range node.Attr {
			if attr.Key == "href" {
				*urlLinks = append(*urlLinks, attr.Val)
			}
		}
	}
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		findLink(child, urlLinks)
	}
}
