package countsAtLeast

import (
	"slices"
	"testing"
)

func TestCountsAtLeast(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected []int
	}{
		{
			name:     "example with duplicates",
			nums:     []int{1, 2, 2, 3},
			expected: []int{0, 4, 3, 1, 0},
		},
		{
			name:     "single element",
			nums:     []int{1},
			expected: []int{0, 1},
		},
		{
			name:     "all values are one",
			nums:     []int{1, 1, 1, 1},
			expected: []int{0, 4, 0, 0, 0},
		},
		{
			name:     "all values are maximum",
			nums:     []int{4, 4, 4, 4},
			expected: []int{0, 4, 4, 4, 4},
		},
		{
			name:     "every possible value once",
			nums:     []int{1, 2, 3, 4, 5},
			expected: []int{0, 5, 4, 3, 2, 1},
		},
		{
			name:     "descending input",
			nums:     []int{5, 4, 3, 2, 1},
			expected: []int{0, 5, 4, 3, 2, 1},
		},
		{
			name:     "gaps between values",
			nums:     []int{1, 1, 4, 4},
			expected: []int{0, 4, 2, 2, 2},
		},
		{
			name:     "mixed values with repeated maximum",
			nums:     []int{2, 5, 3, 5, 1},
			expected: []int{0, 5, 4, 3, 2, 2},
		},
		{
			name:     "no ones",
			nums:     []int{2, 2, 3, 3},
			expected: []int{0, 4, 4, 2, 0},
		},
		{
			name:     "several duplicates at different thresholds",
			nums:     []int{1, 3, 3, 3, 6, 6},
			expected: []int{0, 6, 5, 5, 2, 2, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			original := slices.Clone(tt.nums)

			result := countsAtLeast(tt.nums)

			if !slices.Equal(result, tt.expected) {
				t.Fatalf(
					"countsAtLeast(%v) = %v, expected %v",
					tt.nums,
					result,
					tt.expected,
				)
			}

			if !slices.Equal(tt.nums, original) {
				t.Fatalf(
					"countsAtLeast modified input: got %v, original was %v",
					tt.nums,
					original,
				)
			}
		})
	}
}
