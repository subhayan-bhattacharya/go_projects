package frequencysort

import (
	"slices"
	"testing"
)

func TestFrequencySort(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected []int
	}{
		{
			name:     "different frequencies",
			nums:     []int{1, 1, 2, 2, 2, 3},
			expected: []int{3, 1, 1, 2, 2, 2},
		},
		{
			name:     "same frequency uses descending value",
			nums:     []int{2, 3, 1, 3, 2},
			expected: []int{1, 3, 3, 2, 2},
		},
		{
			name:     "all values unique",
			nums:     []int{5, 1, 3, 2},
			expected: []int{5, 3, 2, 1},
		},
		{
			name:     "all values identical",
			nums:     []int{7, 7, 7},
			expected: []int{7, 7, 7},
		},
		{
			name:     "single value",
			nums:     []int{42},
			expected: []int{42},
		},
		{
			name:     "several values tied on frequency",
			nums:     []int{4, 4, 1, 1, 2, 3, 3},
			expected: []int{2, 4, 4, 3, 3, 1, 1},
		},
		{
			name:     "negative numbers",
			nums:     []int{-1, -1, -2, -2, -3},
			expected: []int{-3, -1, -1, -2, -2},
		},
		{
			name:     "zero positive and negative values",
			nums:     []int{0, 0, -1, 2, 2, 3},
			expected: []int{3, -1, 2, 2, 0, 0},
		},
		{
			name:     "four different frequency levels",
			nums:     []int{1, 1, 1, 2, 2, 3, 4, 4, 4, 4},
			expected: []int{3, 2, 2, 1, 1, 1, 4, 4, 4, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := frequencySort(tt.nums)

			if !slices.Equal(result, tt.expected) {
				t.Errorf("frequencySort(%v) = %v; expected %v",
					tt.nums, result, tt.expected)
			}
		})
	}
}
