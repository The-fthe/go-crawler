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
		html, err := GetHMTL(os.Args[1])
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		fmt.Println(html)
		os.Exit(0)
	}
}
