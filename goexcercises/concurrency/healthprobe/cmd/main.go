package main

import (
	"fmt"
	"healthprobe"
)

func main() {
	urls := []string{
		"https://google.com",
		"https://github.com",
		"https://httpbin.org/status/404",
		"https://thisdomain-does-not-exist.com",
	}
	for _, result := range healthprobe.CheckUrls(urls) {
		fmt.Printf("%+v\n", result)
	}
}
