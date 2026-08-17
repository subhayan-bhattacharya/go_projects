package internal

import (
	"errors"
	"fmt"
	"io"
	"iter"
	"net/http"
	"net/url"
	"slices"
	"strings"
)

var ErrCouldNotDownloadHtml = errors.New("could not download html")
var ErrCouldNotReadResponseBody = errors.New("could not read response body")

func GetUrlData(url string) (error, string, string, string) {
	var output string
	var scheme string
	var hostname string
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("%s: %w", url, ErrCouldNotDownloadHtml), output, scheme, hostname
	}
	scheme = resp.Request.URL.Scheme
	hostname = resp.Request.URL.Host
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return ErrCouldNotReadResponseBody, output, scheme, hostname
	}
	output = string(body)
	return nil, output, scheme, hostname
}

func Map[T, U any](data iter.Seq[T], fn func(T) U) iter.Seq[U] {
	return func(yield func(U) bool) {
		for v := range data {
			if !yield(fn(v)) {
				return
			}
		}
	}
}

func Filter[T any](data iter.Seq[T], fn func(T) bool) iter.Seq[T] {
	return func(yield func(T) bool) {
		for v := range data {
			if fn(v) {
				if !yield(v) {
					return
				}
			}
		}
	}
}

func PreProcessUrls(urls []string, baseUrl *url.URL) []string {
	var isAcceptedUrl func(path string) bool
	isAcceptedUrl = func(path string) bool {
		parsed, err := url.Parse(path)
		if err != nil {
			return false
		}
		if parsed.IsAbs() {
			return (parsed.Scheme == "http" || parsed.Scheme == "https") && parsed.Host == baseUrl.Host
		}
		return parsed.Host == "" && parsed.Path != ""
	}
	var addDomainToPath func(string) string
	addDomainToPath = func(path string) string {
		trimmedPath := strings.TrimRight(path, "/")
		relURL, _ := url.Parse(trimmedPath)
		if !strings.HasPrefix(trimmedPath, baseUrl.Scheme) || !strings.HasPrefix(trimmedPath, baseUrl.Host) {
			return baseUrl.ResolveReference(relURL).String()
		}
		return trimmedPath
	}
	filtered := Filter[string](slices.Values(urls), isAcceptedUrl)
	mapped := Map[string, string](filtered, addDomainToPath)
	return slices.Collect(mapped)
}
