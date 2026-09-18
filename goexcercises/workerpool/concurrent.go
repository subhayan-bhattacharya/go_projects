package main

import (
	"errors"

	"workerpool/db"
)

func MapConcurrent[In, Out any](data []In, workers int, fn func(In) (Out, error)) []Out {
	var output []Out
	return output
}

type Result[T any] struct {
	Index int
	Value T
	Error error
}

func SendData(out chan<- Data, usernames []string) {
	for index, username := range usernames {
		data := Data{
			Index:    index,
			Username: username,
		}
		out <- data
	}
}

func ResultForUsername(repo *db.BoltRepository, dataChannel <-chan Data, resultsChannel chan<- Result[bool]) {
	data := <-dataChannel
	username := data.Username
	if username == "" {
		result := Result[bool]{
			Index: 0,
			Value: false,
			Error: errors.New("empty username"),
		}
		resultsChannel <- result
		return
	}
	if len(username) < 3 {
		result := Result[bool]{
			Index: 0,
			Value: false,
			Error: errors.New("too short a username"),
		}
		resultsChannel <- result
		return
	}
	userDetails, err := repo.GetUser(username)
	if err != nil {
		if errors.Is(err, db.ErrUserNotFound) {
			result := Result[bool]{
				Index: 0,
				Value: false,
				Error: db.ErrUserNotFound,
			}
			resultsChannel <- result
		}
	} else {
		if userDetails.Username == username {
			result := Result[bool]{
				Index: 0,
				Value: true,
				Error: nil,
			}
			resultsChannel <- result
		}
	}
	return
}

type Data struct {
	Index    int
	Username string
}
