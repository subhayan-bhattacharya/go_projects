package main

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type Handler func(any) string

var handlers = map[reflect.Type]Handler{}

func Register[T any](fn func(T) string) {
	var zero T
	t := reflect.TypeOf(zero)
	handlers[t] = func(v any) string {
		return fn(v.(T))
	}
}

func Handle(v any) (string, error) {
	fn, ok := handlers[reflect.TypeOf(v)]
	if !ok {
		return "", errors.New("could not find handler")
	}
	return fn(v), nil
}

func main() {
	Register(func(i int) string { return strconv.Itoa(i) })
	Register(func(i string) string { return strings.ToUpper(i) })
	a1, _ := Handle(20)
	a2, _ := Handle("subhayan")
	a3, err := Handle(struct{ Name string }{Name: "Shaayan"})
	fmt.Println(a1, a2)
	if err != nil {
		fmt.Println(err)
		fmt.Println(a3)
	}
}
