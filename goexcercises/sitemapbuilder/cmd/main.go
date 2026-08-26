package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"os"
	"sitemapbuilder"
)

func main() {
	parseUrl := flag.String("url", "http://gophercises.com", "the url that you want to use the sitemap for.")
	maxDepth := flag.Int("depth", 10, "the maximum depth to which you want to recurse to.")
	flag.Parse()
	data := sitemapbuilder.Bfs(*parseUrl, *maxDepth)
	//for _, d := range data {
	//	fmt.Println(d)
	//}
	toXml := sitemapbuilder.UrlSet{
		Xmlns: sitemapbuilder.Xmlns,
	}
	fmt.Print(xml.Header)
	for _, d := range data {
		toXml.Urls = append(toXml.Urls, sitemapbuilder.Loc{
			Value: d,
		})
	}
	encoder := xml.NewEncoder(os.Stdout)
	encoder.Indent("", "  ")
	if err := encoder.Encode(toXml); err != nil {
		panic(err)
	}
}
