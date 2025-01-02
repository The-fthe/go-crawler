package main

import (
	"fmt"
)

const TOO_MANY_ARGUMENT = "too many arguments provided"
const NO_WEBSITE_PROVIDED = "no website provided"
const START_CRAWL = "starting crawl of:"

func cmd(args []string) string {
	argsProg := args[1:]
	argsLen := len(argsProg)

	if argsLen <= 0 {
		return "no website provided"
	}

	if argsLen > 1 {
		return "too many arguments provided"
	}

	BASE_URL := argsProg[0]
	return fmt.Sprintf("starting crawl of: %s", BASE_URL)
}
