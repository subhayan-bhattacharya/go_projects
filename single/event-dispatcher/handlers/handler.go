package handlers

import (
	events "event-dispatcher/events"
	"fmt"
)

type Row struct {
	Username string
	Email    string
	Password string
}

var Database = map[string]Row{}

type Handler[T any] func(T)

func LoggerHandler(event events.UserSignedUp) {
	fmt.Printf("[LOG] User signed up: %s (%s)\n", event.UserName, event.Email)
}

func DatabaseHandler(event events.UserSignedUp) {
	row := Row{
		Username: event.UserName,
		Email:    event.Email,
		Password: "password",
	}
	Database[event.UserName] = row
}

func RiskyHandler(event events.UserSignedUp) {
	if event.UserName == "admin" {
		panic("Unauthorized username used!")
	}
	fmt.Println("[SECURITY] Username check passed.")
}
