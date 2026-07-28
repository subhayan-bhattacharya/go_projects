package healthprobe

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type Result struct {
	URL        string
	StatusCode int
	Duration   time.Duration
	Error      error
}

const workers = 3

func signalWorkIsDone(wg *sync.WaitGroup, channel chan Result) {
	//The idea is to make sure that it waits for the other goroutines to complete
	//and then close the channel so that the range loop does not hang
	wg.Wait()
	close(channel)
}

func feedJobsChannel(jobsChannel chan<- string, urls []string) {
	for _, url := range urls {
		jobsChannel <- url
	}
	close(jobsChannel) // Signal that there are no more jobs coming
}

func launchWorker(jobsChannel <-chan string, resultsChannel chan<- Result, wg *sync.WaitGroup, ctx context.Context, ticker <-chan time.Time) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		case url, ok := <-jobsChannel:
			if !ok {
				fmt.Println("nothing left to consume...")
				return
			}
			fmt.Printf("consumed url %s now waiting for ticker\n", url)
			select {
			case <-ctx.Done():
				resultsChannel <- Result{
					URL:   url,
					Error: ctx.Err(), // context.Canceled
				}
				return
			case <-ticker:
				healthcheck(ctx, resultsChannel, url)
			}
		}
	}
}

//Go fan-out / fan-in orchestration pattern

func CheckUrls(urls []string, ctx context.Context) []Result {
	if len(urls) == 0 {
		return []Result{}
	}
	var wg sync.WaitGroup
	resultsChannel := make(chan Result)
	jobsChannel := make(chan string, len(urls)) // create a buffered channel to feed the urls instead of
	//launching one go routine per url
	feedJobsChannel(jobsChannel, urls)
	results := make([]Result, 0, len(urls))
	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()
	for _ = range workers {
		wg.Add(1)
		go launchWorker(jobsChannel, resultsChannel, &wg, ctx, ticker.C)
	}
	go signalWorkIsDone(&wg, resultsChannel) // Launch this so that this makes sure the channel is closed
	fmt.Println("Now consuming...")
	for result := range resultsChannel {
		results = append(results, result)
	}
	return results
}

func healthcheck(ctx context.Context, result chan<- Result, url string) {
	//I make this a buffered channel so that when timeout happens, then this function would
	//exit and then no one would be listening on this channel anymore for http results
	httpChannel := make(chan Result, 1)
	timer := time.NewTimer(2 * time.Second)
	defer timer.Stop()
	go func() {
		client := http.Client{}
		req, _ := http.NewRequestWithContext(ctx, "GET", url, nil)
		start := time.Now()
		resp, err := client.Do(req)
		duration := time.Since(start)
		if err != nil {
			httpChannel <- Result{
				URL:      url,
				Duration: duration,
				Error:    err,
			}
			return
		}
		defer resp.Body.Close()
		httpResult := Result{
			URL:        url,
			StatusCode: resp.StatusCode,
			Duration:   duration,
			Error:      err,
		}
		httpChannel <- httpResult
	}()
	select {
	case <-ctx.Done():
		result <- Result{
			URL:   url,
			Error: ctx.Err(), // context.Canceled
		}
		return
	case res := <-httpChannel:
		result <- res
	case <-timer.C:
		fmt.Printf("timed out waiting for response for url %s\n", url)
		err := fmt.Errorf("timeout waiting for response")
		result <- Result{
			URL:      url,
			Duration: 2 * time.Second,
			Error:    err,
		}
	}
}
