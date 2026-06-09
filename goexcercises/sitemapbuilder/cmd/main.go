package main

import (
	"fmt"
	"net/http"
	"sitemapbuilder"
)

func main() {
	resp, err := http.Get("https://bongoutsavdresden.de/")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	data, _ := sitemapbuilder.Parse(resp.Body)
	for _, hrefs := range data {
		fmt.Println(hrefs)
	}
}
