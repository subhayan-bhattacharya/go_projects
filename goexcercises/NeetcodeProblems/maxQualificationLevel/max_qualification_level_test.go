package maxQualificationLevel

import (
	"slices"
	"testing"
)

func TestMaxQualificationLevel(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "standard case",
			nums:     []int{3, 0, 6, 1, 5},
			expected: 3,
		},
		{
			name:     "answer not present in array",
			nums:     []int{1, 1, 1, 4, 7},
			expected: 2,
		},
		{
			name:     "all large values",
			nums:     []int{10, 10, 10, 10},
			expected: 4,
		},
		{
			name:     "all zero",
			nums:     []int{0, 0, 0, 0},
			expected: 0,
		},
		{
			name:     "single qualifying value",
			nums:     []int{100},
			expected: 1,
		},
		{
			name:     "single zero",
			nums:     []int{0},
			expected: 0,
		},
		{
			name:     "duplicates near boundary",
			nums:     []int{2, 2, 2, 2, 10},
			expected: 2,
		},
		{
			name:     "boundary between values",
			nums:     []int{1, 1, 4, 4, 4},
			expected: 3,
		},
		{
			name:     "descending input",
			nums:     []int{8, 7, 6, 5, 4},
			expected: 4,
		},
		{
			name:     "mixed duplicates",
			nums:     []int{0, 2, 2, 3, 3, 5},
			expected: 3,
		},
		{
			name:     "large numbers do not make answer exceed length",
			nums:     []int{100, 200, 300},
			expected: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			original := append([]int(nil), tt.nums...)

			got := maxQualificationLevel(tt.nums)

			if got != tt.expected {
				t.Fatalf(
					"maxQualificationLevel(%v) = %d; expected %d",
					tt.nums,
					got,
					tt.expected,
				)
			}

			if !slices.Equal(tt.nums, original) {
				t.Fatalf(
					"function modified nums: got %v; original %v",
					tt.nums,
					original,
				)
			}
		})
	}
}
