package main

import (
	"flag"
	"fmt"
	"sitemapbuilder"
)

func main() {
	parseUrl := flag.String("url", "http://gophercises.com", "the url that you want to use the sitemap for.")
	maxDepth := flag.Int("depth", 10, "the maximum depth to which you want to recurse to.")
	flag.Parse()
	data := sitemapbuilder.Bfs(*parseUrl, *maxDepth)
	for _, d := range data {
		fmt.Println(d)
	}
}
