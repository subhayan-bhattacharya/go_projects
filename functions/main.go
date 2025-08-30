package main

import (
	"fmt"
	"strings"
)

type StringTransformer func(string) string

func (s StringTransformer) AndThen(other StringTransformer) StringTransformer {
	return func (input string) string {
		return other(s(input))
	}
}

func (s StringTransformer) Apply(inputs []string) []string {
	results := make([]string, 0, len(inputs))
	for _, value := range inputs {
		results = append(results, s(value))
	}
	return results
}

func AddExclamation(input string) string {
	var builder strings.Builder
	builder.WriteString(input)
	builder.WriteString(" !!!")
	return builder.String()
}

func main() {
	inputs := []string {
		"Shaayan is a good guy",
		"Subhayan is a bad guy",
		"Dimpu is a wonderful lady",
	}
	var transformer StringTransformer = strings.ToUpper
	results := transformer.Apply(inputs)
	for _, result := range results {
		fmt.Println(result)
	}

	newResult := transformer.AndThen(AddExclamation).Apply(inputs)
	for _, result := range newResult {
		fmt.Println(result)
	}
}
