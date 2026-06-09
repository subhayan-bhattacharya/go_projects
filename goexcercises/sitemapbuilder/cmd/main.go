package main

import (
	"flag"
	"fmt"
	"net/http"
	"net/url"
)

func main() {
	website := flag.String("website", "https://gophercises.com/demos/cyoa", "The website to scrape")
	flag.Parse()
	resp, err := http.Get(*website)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	reqUrl := resp.Request.URL
	fmt.Println(reqUrl)
	baseUrl := &url.URL{
		Scheme: reqUrl.Scheme,
		Host:   reqUrl.Host,
	}
	base := baseUrl.String()
	fmt.Println(base)
	//data, _ := sitemapbuilder.Parse(resp.Body)
	////io.Copy(os.Stdout, resp.Body)
	//for _, hrefs := range data {
	//	fmt.Println(hrefs)
	//}
}
