package main

import (
	"fmt"
	"net/url"
	"sync"
)

type Config struct {
	pages              map[string]int
	baseURL            *url.URL
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
}

func configure(rawBaseURL string, maxConcurrency int) (*Config, error) {
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		return nil, err
	}
	c := Config{
		pages:              make(map[string]int),
		baseURL:            baseURL,
		mu:                 new(sync.Mutex),
		concurrencyControl: make(chan struct{}, maxConcurrency),
		wg:                 new(sync.WaitGroup),
	}
	return &c, nil
}

// require add `Config.wg.Add(1)` when first run, to prevent negative wg
func (c *Config) crawlPage(rawCurrentURL string) {
	c.concurrencyControl <- struct{}{}
	defer func() {
		<-c.concurrencyControl
		c.wg.Done()
	}()

	parseCurrURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Println("Parse Curr Url error: ", err.Error())
		return
	}

	//skip other domain website
	if c.baseURL.Hostname() != parseCurrURL.Hostname() {
		return
	}

	normalizedURL, err := normalizedURL(rawCurrentURL)
	if err != nil {
		return
	}

	if c.addPageVisist(normalizedURL) {
		return
	}

	fmt.Println("Crawl url: ", rawCurrentURL)

	htmlBody, err := GetHMTL(rawCurrentURL)
	if err != nil {
		fmt.Println("GetHTML error: ", err.Error())
		return
	}

	nextURLs, err := getURLsFromHTML(htmlBody, c.baseURL)
	if err != nil {
		fmt.Println("GetURLs error: ", err.Error())
		return
	}

	for _, nextURL := range nextURLs {
		c.wg.Add(1)
		go c.crawlPage(nextURL)
	}
}

func (c *Config) addPageVisist(normalizedURL string) (isFirst bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	_, visited := c.pages[normalizedURL]

	if visited {
		c.pages[normalizedURL]++
		return visited
	}

	c.pages[normalizedURL] = 1
	return visited
}
