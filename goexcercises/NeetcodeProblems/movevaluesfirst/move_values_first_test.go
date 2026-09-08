package movevaluesfirst

import (
	"slices"
	"testing"
)

func TestMoveValuesFirst(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		first    int
		second   int
		expected []int
	}{
		{
			name:     "both priority values appear multiple times",
			nums:     []int{3, 2, 5, 1, 2, 4, 1, 2},
			first:    2,
			second:   1,
			expected: []int{2, 2, 2, 1, 1, 3, 5, 4},
		},
		{
			name:     "both priority values appear once",
			nums:     []int{4, 1, 3, 2, 5},
			first:    3,
			second:   1,
			expected: []int{3, 1, 4, 2, 5},
		},
		{
			name:     "first priority value does not appear",
			nums:     []int{5, 1, 4, 1, 3},
			first:    2,
			second:   1,
			expected: []int{1, 1, 5, 4, 3},
		},
		{
			name:     "second priority value does not appear",
			nums:     []int{5, 2, 4, 2, 3},
			first:    2,
			second:   1,
			expected: []int{2, 2, 5, 4, 3},
		},
		{
			name:     "neither priority value appears",
			nums:     []int{5, 8, 4, 3},
			first:    2,
			second:   1,
			expected: []int{5, 8, 4, 3},
		},
		{
			name:     "priority values already grouped correctly",
			nums:     []int{2, 2, 1, 1, 7, 5, 4},
			first:    2,
			second:   1,
			expected: []int{2, 2, 1, 1, 7, 5, 4},
		},
		{
			name:     "priority values appear in reverse order",
			nums:     []int{1, 1, 4, 2, 5, 2},
			first:    2,
			second:   1,
			expected: []int{2, 2, 1, 1, 4, 5},
		},
		{
			name:     "negative priority values",
			nums:     []int{3, -1, 4, -2, -1, 5, -2},
			first:    -2,
			second:   -1,
			expected: []int{-2, -2, -1, -1, 3, 4, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := moveValuesFirst(tt.nums, tt.first, tt.second)

			if !slices.Equal(result, tt.expected) {
				t.Errorf(
					"moveValuesFirst(%v, %d, %d) = %v; expected %v",
					tt.nums,
					tt.first,
					tt.second,
					result,
					tt.expected,
				)
			}
		})
	}
}
