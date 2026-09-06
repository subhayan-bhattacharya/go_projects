package finderrornums

import (
	"slices"
	"testing"
)

func TestFindErrorNums(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected []int
	}{
		{
			name:     "first",
			nums:     []int{1, 2, 2, 4},
			expected: []int{2, 3},
		},
		{
			name:     "second",
			nums:     []int{3, 2, 3, 4, 6, 5},
			expected: []int{3, 1},
		},
		{
			name:     "third",
			nums:     []int{2, 2},
			expected: []int{2, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findErrorNums(tt.nums)

			if !slices.Equal(result, tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}
