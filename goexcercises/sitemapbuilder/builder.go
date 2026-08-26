package sitemapbuilder

import (
	"errors"
	"fmt"
	"net/url"
	"sitemapbuilder/internal"
	"strings"
)

const Xmlns = "http://www.sitemaps.org/schemas/sitemap/0.9"

type Loc struct {
	Value string `xml:"loc"`
}

type UrlSet struct {
	Xmlns string `xml:"xmlns,attr"`
	Urls  []Loc  `xml:"url"`
}

func Bfs(website string, depth int) []string {
	seen := make(map[string]struct{})
	var queue map[string]struct{}
	nextQueue := map[string]struct{}{website: {}}
	for i := range 100000 {
		fmt.Println("i=", i)
		queue, nextQueue = nextQueue, make(map[string]struct{})
		if len(queue) == 0 {
			break
		}
		for queuedUrl, _ := range queue {
			if _, ok := seen[queuedUrl]; ok {
				continue
			}
			seen[queuedUrl] = struct{}{}
			links, err := GetEmbeddedLinks(queuedUrl)
			if err != nil {
				fmt.Printf("encountered error for url %s", queuedUrl)
				continue
			}
			for _, link := range links {
				if _, ok := seen[link]; !ok {
					nextQueue[link] = struct{}{}
				}
			}
		}
	}
	var result []string
	for url, _ := range seen {
		result = append(result, url)
	}
	return result
}

func GetEmbeddedLinks(website string) ([]string, error) {
	var values []string
	err, data, scheme, hostname := internal.GetUrlData(website)
	if errors.Is(err, internal.ErrCouldNotDownloadHtml) {
		panic(fmt.Sprintf("could not read html: %v", err))
	} else if errors.Is(err, internal.ErrCouldNotReadResponseBody) {
		panic("unable to read response body")
	}
	links, err := internal.Parse(strings.NewReader(data))
	if err != nil {
		return values, fmt.Errorf("could not parse html: %w", err)
	}
	var allUrls []string
	for _, link := range links {
		allUrls = append(allUrls, link.Href)
	}
	//parsed, err := url.Parse(website)
	//if err != nil {
	//	return values, fmt.Errorf("could not parse url")
	//}
	//baseUrl := &url.URL{
	//	Scheme: parsed.Scheme,
	//	Host:   parsed.Host,
	//}
	baseUrl := &url.URL{
		Scheme: scheme,
		Host:   hostname,
	}
	preProcessedUrls := internal.PreProcessUrls(allUrls, baseUrl)
	values = append(values, preProcessedUrls...)
	return values, nil
}
