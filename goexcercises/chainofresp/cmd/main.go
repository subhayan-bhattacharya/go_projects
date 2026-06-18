package main

import (
	"chainofresp"
	"errors"
	"fmt"
	"net/http"
)

func handler(writer http.ResponseWriter, r *http.Request) {
	ratelimiterHandler := chainofresp.RateLimiter{}
	authorisationHandler := chainofresp.AuthenticationHandler{}
	logger := chainofresp.Logger{}
	logger.SetNext(&authorisationHandler).SetNext(&ratelimiterHandler)
	err := logger.Execute(r)
	if err != nil {
		if pathErr, ok := errors.AsType[*chainofresp.PipelineError](err); ok {
			fmt.Println(pathErr)
			writer.WriteHeader(pathErr.StatusCode)
			writer.Write([]byte(pathErr.Error()))
			return
		}
	}
	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte("all good!"))
}

func main() {
	http.HandleFunc("/", handler)
	_ = http.ListenAndServe(":8080", nil)
}
