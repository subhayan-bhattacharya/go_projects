// Implementation of a queue

package main

import (
	"errors"
	"fmt"
)

type Queue[T comparable] struct {
	items []T
}

func (queue *Queue[T]) add(item T) {
	queue.items = append(queue.items, item)
}

func (queue *Queue[T]) remove() (T, bool) {
	if len(queue.items) == 0 {
		var zero T
		return zero, false
	}
	item := queue.items[0]
	queue.items = queue.items[1:]
	return item, true
}

func (queue *Queue[T]) find(value T) (T, error) {
	var zero T
	if len(queue.items) == 0 {
		return zero, errors.New("Queue is empty")
	}
	for _, item := range queue.items {
		if item == value {
			return value, nil
		}
	}
	return zero, nil
}

func main() {
	var queue Queue[int]
	queue.add(1)
	queue.add(34)
	queue.add(90)
	queue.add(111)
	value, ok := queue.find(420)
	if ok != nil {
		fmt.Println("Value is there : ", value)
	} else {
		fmt.Println("Value is not there")
	}
	for {
		value, isEmpty := queue.remove()
		if isEmpty {
			fmt.Println(value)
		} else {
			fmt.Println("The queue is empty")
			break
		}
	}
}
