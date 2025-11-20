package main

import (
	"testing"
)

type Input struct {
	originalPrice int
	age int
	result float64
}

var testData = []Input {
	Input{
		originalPrice: 30000,
		age: 5,
		result: 21000.00,
	},
	Input{
		originalPrice: 30000,
		age: 11,
		result: 15000.00,
	},
	Input{
		originalPrice: 30000,
		age: 7,
		result: 21000.00,
	},
	Input{
		originalPrice: 30000,
		age: 2,
		result: 24000.00,
	},
}

func TestCalculateResellPrice(t *testing.T) {
	for _, data := range testData {
		actual := CalculateResellPrice(data.originalPrice, data.age)
		if actual != data.result {
			t.Errorf("Expected %f got %f\n", data.result, actual)
		}
	}
}