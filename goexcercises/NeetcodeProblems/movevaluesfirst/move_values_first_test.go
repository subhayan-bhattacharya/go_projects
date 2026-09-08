package movevaluesfirst

import (
	"slices"
	"testing"
)

func TestMoveValueFirst(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected []int
	}{
		{
			name:     "target appears multiple times",
			nums:     []int{3, 2, 5, 2, 4, 2},
			target:   2,
			expected: []int{2, 2, 2, 3, 5, 4},
		},
		{
			name:     "target appears once",
			nums:     []int{4, 1, 3, 2},
			target:   3,
			expected: []int{3, 4, 1, 2},
		},
		{
			name:     "target does not appear",
			nums:     []int{5, 1, 4, 3},
			target:   2,
			expected: []int{5, 1, 4, 3},
		},
		{
			name:     "all values are target",
			nums:     []int{7, 7, 7, 7},
			target:   7,
			expected: []int{7, 7, 7, 7},
		},
		{
			name:     "targets already first",
			nums:     []int{2, 2, 3, 5, 4},
			target:   2,
			expected: []int{2, 2, 3, 5, 4},
		},
		{
			name:     "targets at the end",
			nums:     []int{8, 6, 4, 1, 1},
			target:   1,
			expected: []int{1, 1, 8, 6, 4},
		},
		{
			name:     "negative target",
			nums:     []int{3, -1, 4, -1, 2},
			target:   -1,
			expected: []int{-1, -1, 3, 4, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := moveValueFirst(tt.nums, tt.target)

			if !slices.Equal(result, tt.expected) {
				t.Errorf(
					"moveValueFirst(%v, %d) = %v; expected %v",
					tt.nums,
					tt.target,
					result,
					tt.expected,
				)
			}
		})
	}
}
