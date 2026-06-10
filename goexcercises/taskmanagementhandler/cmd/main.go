package main

import (
	"fmt"
	"net/http"
	"taskmanagementhandler"
)

func main() {
	taskHandler := taskmanagementhandler.RealTaskHandler{}
	loggingHandler := taskmanagementhandler.LoggingDecorator{Handler: taskHandler}
	mux := http.NewServeMux()
	mux.HandleFunc("/handler", loggingHandler.Handle)
	fmt.Println("Server starting on :8080...")
	http.ListenAndServe(":8080", mux)
}
