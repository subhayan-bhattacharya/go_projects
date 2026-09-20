package main

import (
	"cmp"
	"errors"
	"slices"
	"sync"

	"workerpool/db"
)

func MapConcurrent[In, Out any](
	inputData []In, workers int, fn func(data Data[In]) Result[Out],
) []Result[Out] {
	var output []Result[Out]
	dataChannel := make(chan Data[In])
	resultsChannel := make(chan Result[Out])
	go SendData(dataChannel, inputData)
	var processWg sync.WaitGroup
	for range workers {
		processWg.Add(1)
		go processData(&processWg, dataChannel, resultsChannel, fn)
	}
	go func() {
		processWg.Wait()
		close(resultsChannel)
	}()
	for val := range resultsChannel {
		output = append(output, val)
	}
	slices.SortFunc(
		output, func(a, b Result[Out]) int {
			return cmp.Compare(a.Index, b.Index)
		},
	)
	return output
}

type Result[T any] struct {
	Index int
	Value T
	Error error
}

func SendData[In any](out chan<- Data[In], inputs []In) {
	for index, input := range inputs {
		data := Data[In]{
			Index:    index,
			Username: input,
		}
		out <- data
	}
	close(out)
}

func processData[In, Out any](
	wg *sync.WaitGroup, inChannel <-chan Data[In],
	resultsChannel chan<- Result[Out], fn func(Data[In]) Result[Out],
) {
	defer wg.Done()
	for input := range inChannel {
		output := fn(input)
		resultsChannel <- output
	}
}

func ResultForUsername(repo *db.BoltRepository) func(data Data[string]) Result[bool] {
	return func(data Data[string]) Result[bool] {
		username := data.Username
		var result Result[bool]
		if username == "" {
			return Result[bool]{
				Index: data.Index,
				Value: false,
				Error: errors.New("empty username"),
			}
		}
		if len(username) < 3 {
			return Result[bool]{
				Index: data.Index,
				Value: false,
				Error: errors.New("too short a username"),
			}
		}
		userDetails, err := repo.GetUser(username)
		if err != nil {
			if errors.Is(err, db.ErrUserNotFound) {
				result = Result[bool]{
					Index: data.Index,
					Value: false,
					Error: db.ErrUserNotFound,
				}
			}
		} else {
			if userDetails.Username == username {
				result = Result[bool]{
					Index: data.Index,
					Value: true,
					Error: nil,
				}
			}
		}
		return result
	}
}

type Data[In any] struct {
	Index    int
	Username In
}
