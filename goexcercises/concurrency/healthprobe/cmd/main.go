package main

import (
	"fmt"
	"healthprobe"
	"net/http"
	"net/http/httptest"
	"time"
)

func main() {
	slowServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(5 * time.Second) // longer than your 2s timeout
		w.WriteHeader(http.StatusOK)
	}))
	defer slowServer.Close()
	urls := []string{
		"https://www.google.com",               // fast — should succeed quickly
		"https://www.github.com",               // fast — should succeed quickly
		"https://httpbin.org/status/500",       // fast — HTTP 500, no Go error
		slowServer.URL,                         // very slow — should TIMEOUT
		"https://thisurldoesnotexist12345.com", // DNS failure — fast error
	}
	for _, result := range healthprobe.CheckUrls(urls) {
		fmt.Printf("%+v\n", result)
	}
}
