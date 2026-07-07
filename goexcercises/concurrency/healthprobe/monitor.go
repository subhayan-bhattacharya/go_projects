package healthprobe

import (
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

func signalWorkIsDone(wg *sync.WaitGroup, channel chan Result) {
	//The idea is to make sure that it waits for the other goroutines to complete
	//and then close the channel so that the range loop does not hang
	wg.Wait()
	close(channel)
}

//Go fan-out / fan-in orchestration pattern

func CheckUrls(urls []string) []Result {
	if len(urls) == 0 {
		return []Result{}
	}
	var wg sync.WaitGroup // this is used in line number 34 in a separate go routine
	channel := make(chan Result)
	results := make([]Result, 0, len(urls))
	for _, url := range urls {
		wg.Add(1)
		go healthcheck(channel, url, &wg)
	}
	go signalWorkIsDone(&wg, channel) // Launch this so that this makes sure the channel is closed
	fmt.Println("Now consuming...")
	for result := range channel {
		results = append(results, result)
	}
	return results
}

func healthcheck(result chan<- Result, url string, wg *sync.WaitGroup) {
	defer wg.Done()
	//I make this a buffered channel so that when timeout happens, then this function would
	//exit and then no one would be listening on this channel anymore for http results
	httpChannel := make(chan Result, 1)
	timer := time.NewTimer(2 * time.Second)
	defer timer.Stop()
	go func() {
		client := http.Client{}
		req, _ := http.NewRequest("GET", url, nil)
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
