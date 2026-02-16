package main

import (
	"errors"
	"fmt"
	"strings"
)

type Converter func(string) (string, error)

var Converters = map[string]Converter{}

func RegisterConverter(name string, fn Converter) error {
	convertedName := strings.ToLower(name)
	if convertedName == "" {
		return errors.New("cannot supply an empty name")
	}
	if fn == nil {
		return errors.New("the function cannot be nil")
	}
	_, ok := Converters[convertedName]
	if ok {
		return errors.New("the converter exists")
	}
	Converters[convertedName] = fn
	return nil
}

func Convert(value string, converter string) (string, error) {
	fn, ok := Converters[converter]
	if !ok {
		return "", errors.New("unable to find a converter")
	}
	return fn(value)
}

func ConvertToLower(input string) (string, error) {
	return strings.ToLower(input), nil
}

func TrimSpace(input string) (string, error) {
	return strings.TrimSpace(input), nil
}

func init() {
	_ = RegisterConverter("lower", ConvertToLower)
	_ = RegisterConverter("trimmer", TrimSpace)
}

func main() {
	out, _ := Convert("upper", "lower")
	fmt.Println(out)

	out, _ = Convert("trim  ", "trimmer")
	fmt.Println(out)
}
