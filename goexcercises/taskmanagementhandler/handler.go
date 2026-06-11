package taskmanagementhandler

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"time"

	"golang.org/x/time/rate"
)

type Task struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

type TaskHandler interface {
	Handle(w http.ResponseWriter, r *http.Request)
}

type RealTaskHandler struct {
}

func (handler RealTaskHandler) Handle(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	task := Task{
		Id:    42,
		Title: "Learn decorators",
		Done:  false,
	}
	err := json.NewEncoder(w).Encode(task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (r *responseWriter) WriteHeader(statusCode int) {
	r.statusCode = statusCode
	r.ResponseWriter.WriteHeader(statusCode)
}

type RateLimitDecorator struct {
	Handler      TaskHandler
	LimitsByHost map[string]*rate.Limiter
}

func (r *RateLimitDecorator) Handle(writer http.ResponseWriter, request *http.Request) {
	host := request.Host
	slog.Info("Checking limit for host : ", "host", host)
	limiter, ok := r.LimitsByHost[host]
	if ok {
		if !limiter.Allow() {
			slog.Warn("rate limit exceeded for host : ", "host", host)
			http.Error(writer, http.StatusText(http.StatusTooManyRequests), http.StatusTooManyRequests)
			return
		}
	} else {
		r.LimitsByHost[host] = rate.NewLimiter(rate.Every(20*time.Second), 3)
	}

	r.Handler.Handle(writer, request)
}

type LoggingDecorator struct {
	Handler TaskHandler
}

func (l LoggingDecorator) Handle(w http.ResponseWriter, r *http.Request) {
	slog.Info("Calling task handler", "Path: ", r.URL.Path, "Method: ", r.Method)
	//rec := httptest.NewRecorder()
	responseWriter := responseWriter{
		ResponseWriter: w,
		statusCode:     0,
	}
	l.Handler.Handle(&responseWriter, r)
	slog.Info("Response from handler: ", "Status code: ", responseWriter.statusCode)
	//w.Header().Set("Content-Type", rec.Header().Get("Content-Type"))
	//w.WriteHeader(rec.Code)
	//w.Write([]byte(rec.Body.String()))
}
