package taskmanagementhandler

import (
	"encoding/json"
	"log/slog"
	"net/http"
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
