package main

import (
	"fmt"
	"os"
)

func main() {
	baseURLRaw, maxConcurreny, maxPage, err := cmd(os.Args)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(START_CRAWL, " ", baseURLRaw)

	rawUrl := os.Args[1]

	c, err := configure(rawUrl, maxConcurreny, maxPage)
	if err != nil {
		fmt.Println("", err.Error())
		os.Exit(1)
	}

	c.wg.Add(1)
	c.crawlPage(rawUrl)
	c.wg.Wait()

	fmt.Println("MaxPage Count: ", c.CurrPageCount())
	for link, count := range c.pages {
		fmt.Println(link, ": ", count)
	}
	os.Exit(0)
}
