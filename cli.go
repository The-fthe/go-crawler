package main

import (
	"fmt"
	"net/url"
	"strconv"
)

const TOO_MANY_ARGUMENT = "too many arguments provided"
const NO_WEBSITE_PROVIDED = "no website provided"
const START_CRAWL = "starting crawl of:"

func cmd(args []string) (string, int, int, error) {
	argsProg := args[1:]
	argsLen := len(argsProg)

	if argsLen < 1 {
		return "", 0, 0, fmt.Errorf("no website provided")
	}
	if argsLen < 2 {
		return "", 0, 0, fmt.Errorf("no maxConcurrency provided")
	}
	if argsLen < 3 {
		return "", 0, 0, fmt.Errorf("no maxPages provided")
	}

	if argsLen > 3 {
		return "", 0, 0, fmt.Errorf("too many arguments provided")
	}

	baseURLRaw := argsProg[0]
	println("baseURLRaw: ", baseURLRaw)
	println("maxConcurrency: ", argsProg[1])
	println("MaxPage: ", argsProg[2])
	baseURL, err := url.Parse(baseURLRaw)
	if err != nil {
		return "", 0, 0, fmt.Errorf("BaseURL parse failed: %v", err)
	}

	maxConcurrency, err := strconv.Atoi(argsProg[1])
	if err != nil {
		return "", 0, 0, fmt.Errorf("maxConcurrency convert failed: %v", err)
	}

	maxPage, err := strconv.Atoi(argsProg[2])
	if err != nil {
		return "", 0, 0, fmt.Errorf("MaxPage convert failed: %v", err)
	}

	return baseURL.String(), maxConcurrency, maxPage, nil
}
