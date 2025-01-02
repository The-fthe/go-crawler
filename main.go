package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	result := cmd(os.Args)
	if result == NO_WEBSITE_PROVIDED {
		fmt.Println(result)
		os.Exit(1)
	} else if result == TOO_MANY_ARGUMENT {
		fmt.Println(result)
		os.Exit(1)
	} else if strings.Contains(result, START_CRAWL) {
		fmt.Println(result)

		rawUrl := os.Args[1]
		pages := make(map[string]int)
		crawlPage(rawUrl, rawUrl, pages)

		for link, count := range pages {
			fmt.Println(link, ": ", count)
		}
	}
}
