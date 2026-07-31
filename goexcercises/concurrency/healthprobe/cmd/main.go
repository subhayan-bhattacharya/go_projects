package main

import (
	"context"
	"fmt"
	"healthprobe"
	"net/http"
	"net/http/httptest"
	"os"
	"os/signal"
	"sync"
	"time"
)

type Stats struct {
	Checks       int
	Failures     int
	TotalLatency time.Duration
}

type SafeCounter struct {
	mu   sync.Mutex
	data map[string]Stats
}

func updateStats(statsChannel <-chan healthprobe.Result, counter *SafeCounter) {
	for result := range statsChannel {
		counter.mu.Lock()
		stats, ok := counter.data[result.URL]
		if ok {
			stats.Checks += 1
			if result.Error != nil {
				stats.Failures += 1
			}
			totalLatency := stats.TotalLatency + result.Duration
			stats.TotalLatency = totalLatency
		} else {
			failure := 0
			if result.Error != nil {
				failure = 1
			}
			stats = Stats{
				Checks:       1,
				Failures:     failure,
				TotalLatency: result.Duration,
			}
		}
		counter.data[result.URL] = stats
		counter.mu.Unlock()
	}
}

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
	store := &SafeCounter{
		data: make(map[string]Stats),
	}
	urls := []string{
		"https://www.google.com",               // fast — should succeed quickly
		"https://www.github.com",               // fast — should succeed quickly
		"https://httpbin.org/status/500",       // fast — HTTP 500, no Go error
		slowServer.URL,                         // very slow — should TIMEOUT
		"https://thisurldoesnotexist12345.com", // DNS failure — fast error
		slowServer.URL,
		slowServer.URL,
	}
	resultsChannel := make(chan healthprobe.Result)
	statsChannel := make(chan healthprobe.Result)
	go updateStats(statsChannel, store)
	go healthprobe.Broadcast(resultsChannel, statsChannel)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		healthprobe.CheckUrls(urls, ctx, resultsChannel)
	}()

	if ctx.Err() != nil {
		fmt.Println("context was cancelled mid flight...")
		store.mu.Lock()
		fmt.Println(store.data)
		store.mu.Unlock()
		fmt.Println()
		return
	}
	orchestrate(ticker, urls, ctx, resultsChannel, &wg)
	wg.Wait()
	close(resultsChannel)
	store.mu.Lock()
	fmt.Println(store.data)
	store.mu.Unlock()
}

func orchestrate(ticker *time.Ticker, urls []string, ctx context.Context, resultsChannel chan healthprobe.Result, wg *sync.WaitGroup) {
	for {
		select {
		case <-ticker.C:
			wg.Add(1)
			go func() {
				defer wg.Done()
				healthprobe.CheckUrls(urls, ctx, resultsChannel)
			}()
		case <-ctx.Done():
			fmt.Println("interrupted due to context cancellation..exiting orchestrate function")
			return
		}
	}
}
