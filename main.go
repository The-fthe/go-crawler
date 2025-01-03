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

		c, err := configure(rawUrl, 1)
		if err != nil {
			fmt.Println("Configure crawler error: ", err.Error())
			os.Exit(1)
		}

		c.crawlPage(rawUrl)

		c.wg.Wait()
		for link, count := range c.pages {
			fmt.Println(link, ": ", count)
		}
	}
}
