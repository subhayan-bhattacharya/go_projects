package healthprobe

import (
	"net/http"
	"net/http/httptest"
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
	results := CheckUrls(urls)
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
	results := CheckUrls([]string{})
	if len(results) != 0 {
		t.Errorf("expected 0 results for empty list, got %d", len(results))
	}
}
