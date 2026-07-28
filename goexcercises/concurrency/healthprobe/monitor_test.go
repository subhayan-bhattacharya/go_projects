package healthprobe

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"slices"
	"sync"
	"testing"
	"time"
)

func TestCheckUrls(t *testing.T) {
	okServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer okServer.Close()

	errorServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer errorServer.Close()

	slowServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(2 * time.Second)
		w.WriteHeader(http.StatusOK)
	}))
	defer slowServer.Close()

	urls := []string{
		okServer.URL,
		errorServer.URL,
		slowServer.URL,
		"http://this-domain-definitely-does-not-exist.local",
	}

	start := time.Now()
	results := CheckUrls(urls, context.Background())
	duration := time.Since(start)

	if len(results) != len(urls) {
		t.Fatalf("expected %d results, got %d", len(urls), len(results))
	}

	if duration > 6*time.Second {
		t.Errorf("execution took too long (%v), concurrency might be broken", duration)
	}

	foundOk := false
	foundErrStatus := false
	foundDialErr := false

	for _, res := range results {
		if res.URL == okServer.URL && res.StatusCode == http.StatusOK {
			foundOk = true
		}
		if res.URL == errorServer.URL && res.StatusCode == http.StatusInternalServerError {
			foundErrStatus = true
		}
		if res.URL == "http://this-domain-definitely-does-not-exist.local" && res.Error != nil {
			foundDialErr = true
		}

		if res.URL == slowServer.URL {
			if res.Error == nil || res.Error.Error() != "timeout waiting for response" {
				t.Errorf("expected timeout error, got %v", res.Error)
			}
		}
	}

	if !foundOk {
		t.Error("did not find successful result for okServer")
	}
	if !foundErrStatus {
		t.Error("did not find 500 status result for errorServer")
	}
	if !foundDialErr {
		t.Error("did not find dial error for non-existent domain")
	}
}

func TestCheckUrls_EmptyList(t *testing.T) {
	results := CheckUrls([]string{}, context.Background())
	if len(results) != 0 {
		t.Errorf("expected 0 results for empty list, got %d", len(results))
	}
}

func TestCheckUrls_RateLimitsRequestStarts(t *testing.T) {
	var mu sync.Mutex
	requestTimes := make([]time.Time, 0, 3)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		requestTimes = append(requestTimes, time.Now())
		mu.Unlock()

		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	urls := []string{
		server.URL,
		server.URL,
		server.URL,
	}

	results := CheckUrls(urls, context.Background())
	if len(results) != len(urls) {
		t.Fatalf("expected %d results, got %d", len(urls), len(results))
	}

	mu.Lock()
	times := slices.Clone(requestTimes)
	mu.Unlock()

	if len(times) != len(urls) {
		t.Fatalf("expected the server to receive %d requests, got %d", len(urls), len(times))
	}

	const minimumGap = 400 * time.Millisecond
	for i := 1; i < len(times); i++ {
		gap := times[i].Sub(times[i-1])
		if gap < minimumGap {
			t.Errorf(
				"requests %d and %d started only %v apart; expected at least %v",
				i,
				i+1,
				gap,
				minimumGap,
			)
		}
	}
}

func TestCheckUrls_ContextCancelledWhileWaitingForRateLimit(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	// More URLs than workers so some jobs are still queued when we cancel.
	urls := []string{
		server.URL,
		server.URL,
		server.URL,
		server.URL,
		server.URL,
		server.URL,
	}

	go func() {
		time.Sleep(100 * time.Millisecond)
		cancel() // cancels the context.
	}()

	start := time.Now()
	results := CheckUrls(urls, ctx)
	elapsed := time.Since(start)

	if elapsed > 3*time.Second {
		t.Errorf("expected fast shutdown, took %v", elapsed)
	}

	if len(results) == 0 {
		t.Error("expected at least some partial results")
	}

	var foundCancel bool
	for _, res := range results {
		if errors.Is(res.Error, context.Canceled) {
			foundCancel = true
			break
		}
	}
	if !foundCancel {
		t.Errorf("expected at least one context.Canceled result, got %v", results)
	}
}
