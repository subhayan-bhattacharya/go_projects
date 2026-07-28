package main

import (
	"context"
	"fmt"
	"healthprobe"
	"net/http"
	"net/http/httptest"
	"os"
	"os/signal"
	"time"
)

func main() {
	slowServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(5 * time.Second) // longer than your 2s timeout
		w.WriteHeader(http.StatusOK)
	}))
	defer slowServer.Close()
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	urls := []string{
		"https://www.google.com",               // fast — should succeed quickly
		"https://www.github.com",               // fast — should succeed quickly
		"https://httpbin.org/status/500",       // fast — HTTP 500, no Go error
		slowServer.URL,                         // very slow — should TIMEOUT
		"https://thisurldoesnotexist12345.com", // DNS failure — fast error
		slowServer.URL,
		slowServer.URL,
	}
	for _, result := range healthprobe.CheckUrls(urls, ctx) {
		fmt.Printf("%+v\n", result)
	}
	if ctx.Err() != nil {
		fmt.Println("context was cancelled mid flight...")
		return
	}
	summary := make([]healthprobe.Result, 0, len(urls))
	for {
		select {
		case <-ticker.C:
			summary = nil
			for index, result := range healthprobe.CheckUrls(urls, ctx) {
				fmt.Printf("Result number: %d %+v\n", index, result)
				summary = append(summary, result)
			}
		case <-ctx.Done():
			fmt.Println("interrupted due to context cancellation")
			fmt.Printf("the last round fetched %d results \n", len(summary))
			return
		}
	}
}
