package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"workerpool/db"

	"github.com/brianvoe/gofakeit/v7"
)

func MapConcurrent[In, Out any](data []In, workers int, fn func(In) (Out, error)) []Out {
	var output []Out
	return output
}

func seedDb(repo *db.BoltRepository) {
	for _ = range 100 {
		user := db.User{
			Username:  gofakeit.Username(),
			Email:     gofakeit.Email(),
			FirstName: gofakeit.FirstName(),
			LastName:  gofakeit.LastName(),
		}
		_ = repo.AddUser(user)
	}
}

type Result[T any] struct {
	Index int
	Value T
	Error error
}

func resultForUsername(repo *db.BoltRepository, username string) Result[bool] {
	var result Result[bool]
	if username == "" {
		return Result[bool]{
			Index: 0,
			Value: false,
			Error: errors.New("empty username"),
		}
	}
	if len(username) < 3 {
		return Result[bool]{
			Index: 0,
			Value: false,
			Error: errors.New("too short a username"),
		}
	}
	userDetails, err := repo.GetUser(username)
	if err != nil {
		if errors.Is(err, db.ErrUserNotFound) {
			return Result[bool]{
				Index: 0,
				Value: false,
				Error: db.ErrUserNotFound,
			}
		}
	} else {
		if userDetails.Username == username {
			return Result[bool]{
				Index: 0,
				Value: true,
				Error: nil,
			}
		}
	}
	return result
}

func main() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	dbPath := filepath.Join(dir, "users.db")
	// Initialize repository with dependency injection
	repo, err := db.NewBoltRepository(dbPath)
	must(err)
	defer repo.Close()
	seedDb(repo)
}

func must(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
