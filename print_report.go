package main

import (
	"fmt"
	"sort"
)

type Page struct {
	URL   string
	Count int
}

func sortPages(pages map[string]int) []Page {
	sortedPages := []Page{}
	for link, count := range pages {
		sortedPages = append(sortedPages, Page{URL: link, Count: count})
	}
	sort.Slice(sortedPages, func(i, j int) bool {
		return sortedPages[i].URL < sortedPages[j].URL
	})

	sort.Slice(sortedPages, func(i, j int) bool {
		return sortedPages[i].Count > sortedPages[j].Count
	})
	return sortedPages
}

func (c *Config) printReport() {
	fmt.Println("=============================")
	fmt.Printf("  REPORT for %s\n", c.baseURL.String())
	fmt.Println("=============================")

	sortedPages := sortPages(c.pages)
	for _, linkCount := range sortedPages {
		fmt.Printf("Found %d internal links to %s\n", linkCount.Count, linkCount.URL)
	}
}
