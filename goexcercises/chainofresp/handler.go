package chainofresp

import (
	"errors"
	"fmt"
	"net/http"
)

type Handler interface {
	Execute(request *http.Request) *PipelineError
	SetNext(handler Handler) Handler
}

type PipelineError struct {
	err        error
	StatusCode int
}

func (p *PipelineError) Error() string {
	return fmt.Sprintf("status %d and error %v\n", p.StatusCode, p.err)
}

type BaseHandler struct {
	next Handler
}

func (b *BaseHandler) ExecuteNext(request *http.Request) *PipelineError {
	if b.next != nil {
		return b.next.Execute(request)
	}
	return nil
}

func (b *BaseHandler) SetNext(handler Handler) Handler {
	b.next = handler
	return handler
}

type Logger struct {
	BaseHandler
}

func (l *Logger) Execute(request *http.Request) *PipelineError {
	verb := request.Method
	uri := request.RequestURI
	fmt.Printf("the incoming request has %s as verb and uri is %s\n", verb, uri)
	return l.ExecuteNext(request)
}

type AuthenticationHandler struct {
	BaseHandler
}

func (a *AuthenticationHandler) Execute(request *http.Request) *PipelineError {
	token := request.Header.Get("Authorization")
	if token == "" {
		return &PipelineError{
			err:        errors.New("no token found inside authorisation header"),
			StatusCode: http.StatusUnauthorized,
		}
	}
	return a.ExecuteNext(request)
}

type RateLimiter struct {
	BaseHandler
}

func (r *RateLimiter) Execute(request *http.Request) *PipelineError {
	value := request.Header.Get("X-Block-Me")
	if value == "true" {
		return &PipelineError{
			err:        errors.New("too many requests"),
			StatusCode: http.StatusTooManyRequests,
		}
	}
	return r.ExecuteNext(request)
}
