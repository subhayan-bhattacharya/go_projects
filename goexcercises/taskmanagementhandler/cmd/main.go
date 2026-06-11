package main

import (
	"fmt"
	"net/http"
	"taskmanagementhandler"

	"golang.org/x/time/rate"
)

func main() {
	taskHandler := taskmanagementhandler.RealTaskHandler{}
	rateLimitingHander := taskmanagementhandler.RateLimitDecorator{
		Handler:      taskHandler,
		LimitsByHost: make(map[string]*rate.Limiter),
	}
	loggingHandler := taskmanagementhandler.LoggingDecorator{Handler: &rateLimitingHander}
	mux := http.NewServeMux()
	mux.HandleFunc("/handler", loggingHandler.Handle)
	fmt.Println("Server starting on :8080...")
	http.ListenAndServe(":8080", mux)
}
