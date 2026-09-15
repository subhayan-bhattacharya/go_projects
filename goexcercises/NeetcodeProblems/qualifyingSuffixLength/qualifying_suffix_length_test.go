package qualifyingSuffixLength

import (
	"slices"
	"testing"
)

func TestQualifyingSuffixLength(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		minimum  int
		expected int
	}{
		{
			name:     "boundary in middle",
			nums:     []int{7, 2, 5, 1, 9},
			minimum:  5,
			expected: 3,
		},
		{
			name:     "duplicate values exactly at boundary",
			nums:     []int{6, 2, 4, 4, 9, 4},
			minimum:  4,
			expected: 5,
		},
		{
			name:     "only largest value qualifies",
			nums:     []int{2, 8, 3, 1},
			minimum:  8,
			expected: 1,
		},
		{
			name:     "every value qualifies",
			nums:     []int{5, 7, 6, 8},
			minimum:  5,
			expected: 4,
		},
		{
			name:     "no value qualifies",
			nums:     []int{1, 2, 3, 4},
			minimum:  10,
			expected: 0,
		},
		{
			name:     "minimum falls between existing values",
			nums:     []int{1, 3, 7, 9, 12},
			minimum:  6,
			expected: 3,
		},
		{
			name:     "many duplicates below boundary",
			nums:     []int{2, 2, 2, 2, 7, 8},
			minimum:  7,
			expected: 2,
		},
		{
			name:     "many duplicates at boundary",
			nums:     []int{1, 5, 5, 5, 5},
			minimum:  5,
			expected: 4,
		},
		{
			name:     "all values equal and qualify",
			nums:     []int{4, 4, 4, 4},
			minimum:  4,
			expected: 4,
		},
		{
			name:     "all values equal and do not qualify",
			nums:     []int{4, 4, 4, 4},
			minimum:  5,
			expected: 0,
		},
		{
			name:     "single value qualifies",
			nums:     []int{4},
			minimum:  4,
			expected: 1,
		},
		{
			name:     "single value does not qualify",
			nums:     []int{4},
			minimum:  5,
			expected: 0,
		},
		{
			name:     "negative values with negative minimum",
			nums:     []int{-5, -1, -3, 2, 0},
			minimum:  -1,
			expected: 3,
		},
		{
			name:     "minimum smaller than every value",
			nums:     []int{3, 1, 2},
			minimum:  0,
			expected: 3,
		},
		{
			name:     "boundary occurs before repeated larger values",
			nums:     []int{9, 9, 3, 9, 2, 5},
			minimum:  5,
			expected: 4,
		},
		{
			name:     "unsorted input with repeated values on both sides",
			nums:     []int{6, 2, 6, 3, 2, 6, 4},
			minimum:  4,
			expected: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			original := append([]int(nil), tt.nums...)

			got := qualifyingSuffixLength(tt.nums, tt.minimum)

			if got != tt.expected {
				t.Fatalf(
					"qualifyingSuffixLength(%v, %d) = %d; expected %d",
					tt.nums,
					tt.minimum,
					got,
					tt.expected,
				)
			}

			if !slices.Equal(tt.nums, original) {
				t.Fatalf(
					"function modified input: got %v; original was %v",
					tt.nums,
					original,
				)
			}
		})
	}
}
