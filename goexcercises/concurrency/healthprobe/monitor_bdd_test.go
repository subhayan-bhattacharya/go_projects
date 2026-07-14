package healthprobe_test

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"time"

	"healthprobe"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("CheckUrls", func() {
	Context("when given multiple URLs", Ordered, func() {
		var (
			okServer    *httptest.Server
			errorServer *httptest.Server
			slowServer  *httptest.Server
			urls        []string
			results     []healthprobe.Result
			duration    time.Duration
			badDomain   = "http://this-domain-definitely-does-not-exist.local"
		)

		BeforeAll(func() {
			okServer = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
			}))

			errorServer = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusInternalServerError)
			}))

			slowServer = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				time.Sleep(2 * time.Second)
				w.WriteHeader(http.StatusOK)
			}))

			urls = []string{
				okServer.URL,
				errorServer.URL,
				slowServer.URL,
				badDomain,
			}

			start := time.Now()
			results = healthprobe.CheckUrls(urls, context.Background())
			duration = time.Since(start)
		})

		AfterAll(func() {
			okServer.Close()
			errorServer.Close()
			slowServer.Close()
		})

		It("returns one result per URL", func() {
			Expect(results).To(HaveLen(len(urls)))
		})

		It("completes in roughly the slowest request time, not the sum", func() {
			Expect(duration).To(BeNumerically("<", 6*time.Second))
		})

		It("reports a successful 200 response", func() {
			var found bool
			for _, res := range results {
				if res.URL == okServer.URL && res.StatusCode == http.StatusOK && res.Error == nil {
					found = true
					break
				}
			}
			Expect(found).To(BeTrue(), "expected a 200 OK result for the healthy server")
		})

		It("reports a non-2xx HTTP status without a Go error", func() {
			var found bool
			for _, res := range results {
				if res.URL == errorServer.URL && res.StatusCode == http.StatusInternalServerError && res.Error == nil {
					found = true
					break
				}
			}
			Expect(found).To(BeTrue(), "expected a 500 status result for the error server")
		})

		It("reports a dial error for an unreachable domain", func() {
			var found bool
			for _, res := range results {
				if res.URL == badDomain && res.Error != nil {
					found = true
					break
				}
			}
			Expect(found).To(BeTrue(), "expected a dial error for the unreachable domain")
		})

		It("reports a timeout error for a slow URL", func() {
			var slowResult healthprobe.Result
			var found bool
			for _, res := range results {
				if res.URL == slowServer.URL {
					slowResult = res
					found = true
					break
				}
			}
			Expect(found).To(BeTrue(), "expected a result for the slow server")
			Expect(slowResult.Error).To(MatchError("timeout waiting for response"))
		})
	})

	Context("when given an empty list", func() {
		It("returns an empty slice without blocking", func() {
			results := healthprobe.CheckUrls([]string{}, context.Background())
			Expect(results).To(BeEmpty())
		})
	})

	Context("when context is cancelled", Ordered, func() {
		var (
			slowServer *httptest.Server
			fastServer *httptest.Server
			urls       []string
			results    []healthprobe.Result
			elapsed    time.Duration
		)

		BeforeAll(func() {
			slowServer = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				select {
				case <-r.Context().Done():
					return
				case <-time.After(10 * time.Second):
					w.WriteHeader(http.StatusOK)
				}
			}))

			fastServer = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
			}))

			urls = []string{
				fastServer.URL,
				slowServer.URL,
				slowServer.URL,
				slowServer.URL,
				slowServer.URL,
				slowServer.URL,
			}

			ctx, cancel := context.WithCancel(context.Background())
			go func() {
				time.Sleep(100 * time.Millisecond)
				cancel()
			}()

			start := time.Now()
			results = healthprobe.CheckUrls(urls, ctx)
			elapsed = time.Since(start)
		})

		AfterAll(func() {
			slowServer.Close()
			fastServer.Close()
		})

		It("shuts down quickly without waiting for slow servers", func() {
			Expect(elapsed).To(BeNumerically("<", 3*time.Second))
		})

		It("returns at least some results", func() {
			Expect(results).NotTo(BeEmpty())
		})

		It("includes at least one context.Canceled error", func() {
			found := false
			for _, res := range results {
				if errors.Is(res.Error, context.Canceled) {
					found = true
					break
				}
			}
			Expect(found).To(BeTrue(), "expected at least one context.Canceled result")
		})
	})
})
