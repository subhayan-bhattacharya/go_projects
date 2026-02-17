//https://chatgpt.com/c/6992cd90-4064-8388-8db3-23184432c475

package main

import (
	"errors"
	"fmt"
)

type User struct {
	Id    int
	Name  string
	Email string
}

type Order struct {
	Number string
	Total  int
}

type Source[T any] func([]byte) (T, error)
type Transform[A, B any] func(A) (B, error)
type Sink[T any] func(T) error

var sourceReg = map[string]func([]byte) (any, error){}
var transformReg = map[string]func(any) (any, error){}
var sinkReg = map[string]func(any) error{}

func RegisterSource[T any](name string, fn Source[T]) {
	sourceReg[name] = func(data []byte) (any, error) {
		return fn(data)
	}
}

func RegisterTransform[A, B any](name string, fn Transform[A, B]) {
	transformReg[name] = func(input any) (any, error) {
		a, ok := input.(A)
		if !ok {
			return nil, errors.New("wrong type supplied")
		}
		return fn(a)
	}
}

func RegisterSink[T any](input string, fn Sink[T]) {
	sinkReg[input] = func(input any) error {
		t, ok := input.(T)
		if !ok {
			return errors.New("wrong type supplied")
		}
		return fn(t)
	}
}

func RunPipeline[A, B any](sourceName, transformName, sinkName string, data []byte) error {
	src, ok := sourceReg[sourceName]
	if !ok {
		return errors.New("fucked up no function for this source")
	}
	tr, ok := transformReg[transformName]
	if !ok {
		return errors.New("fucked up no function for this transformer")
	}
	sink, ok := sinkReg[sinkName]
	if !ok {
		return errors.New("fucked up no function for this sink")
	}
	v, err := src(data)
	if err != nil {
		return fmt.Errorf("source %q failed: %w", sourceName, err)
	}
	a, ok := v.(A)
	if !ok {
		var zeroA A
		return fmt.Errorf("source returned %T and not %T", a, zeroA)
	}

	t, err := tr(a)
	if err != nil {
		return errors.New("transformation failed")
	}

	b, ok := t.(B)
	if !ok {
		var zeroB B
		return fmt.Errorf("source returned %T and not %T", b, zeroB)
	}
	return sink(b)
}

func main() {

}
