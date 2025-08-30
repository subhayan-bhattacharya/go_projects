package main

import (
	"slices"
	"strings"
	"testing"
)

type InputToOutput struct {
	input string
	output string
}


func testData() []InputToOutput {
	testData := []InputToOutput{
		{
      input: "Subhayan is a great guy",
			output: "SUBHAYAN IS A GREAT GUY",
		},
		{
			input: "Dimpu is a wonderful lady",
			output: "DIMPU IS A WONDERFUL LADY",
		},
	}
	return testData
}

func TestStringTransformerApply(t *testing.T) {
	var transformer StringTransformer = strings.ToUpper
	data := testData()
	inputs := make([]string, 0, len(data))
	results := make([]string, 0, len(data))
	for _, data := range data {
		inputs = append(inputs, data.input)
		results = append(results, data.output)
	}
	actualResults := transformer.Apply(inputs)
	if !slices.Equal(results, actualResults) {
		t.Error("The apply function did not work as expected")
	}
}

func TestStringTransformerAndThen(t *testing.T) {
	var transformer StringTransformer = strings.ToUpper
	data := testData()
	inputs := make([]string, 0, len(data))
	results := make([]string, 0, len(data))
	for _, data := range data {
		inputs = append(inputs, data.input)
		var builder strings.Builder
		builder.WriteString(data.output)
		builder.WriteString(" !!!")
		results = append(results, builder.String())
	}
	actualResults := transformer.AndThen(AddExclamation).Apply(inputs)
	if !slices.Equal(actualResults, results) {
		t.Error("The AndThen function did not work")
	}
}



