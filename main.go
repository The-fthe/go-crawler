package main

import (
	"fmt"
	"net/url"
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

		pages := make(map[string]int)
		rawUrl := os.Args[1]
		baseUrl, err := url.Parse(rawUrl)
		if err != nil {
			fmt.Println("error: ", err.Error())
			os.Exit(1)
		}
		baseUrl.Path = ""
		fmt.Println("baseURl: ", baseUrl)

		crawlPage(baseUrl.String(), rawUrl, pages)
		for link, count := range pages {
			fmt.Println("Link: ", link, ": ", count)
		}
	}
}
