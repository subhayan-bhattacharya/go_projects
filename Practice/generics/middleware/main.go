package main

import (
	"context"
	"fmt"
)

type TaskFunc[T any] func(ctx context.Context, data T) error

type Middleware[T any] func(next TaskFunc[T]) TaskFunc[T]

func Wrap[T any](taskFunc TaskFunc[T], middlewares ...Middleware[T]) TaskFunc[T] {
	for i := len(middlewares) - 1; i >= 0; i-- {
		taskFunc = middlewares[i](taskFunc)
	}
	return taskFunc
}

type Email struct {
	To   string
	Body string
}

func ValidateEmailMiddleware(next TaskFunc[Email]) TaskFunc[Email] {
	return func(ctx context.Context, data Email) error {
		if data.To == "" {
			return fmt.Errorf("The to field cannot be empty in an email")
		}
		if data.Body == "" {
			return fmt.Errorf("The body field cannot be empty in an email")
		}
		return next(ctx, data)
	}
}

func LogMiddleware(next TaskFunc[Email]) TaskFunc[Email] {
	return func(ctx context.Context, data Email) error {
		fmt.Printf("LOG: Dispatching task with data: %v\n", data)
		return next(ctx, data)
	}
}

func SendEmail(ctx context.Context, data Email) error {
	fmt.Printf("Sending email to %s with body %s\n", data.To, data.Body)
	return nil
}

func main() {
	ctx := context.Background()
	chain := Wrap(SendEmail, ValidateEmailMiddleware, LogMiddleware)
	chain(ctx, Email{
		To:   "subhayan.here@gmail.com",
		Body: "How long do you need to learn Go ??",
	})
}
