package moveValuesByOrder

import (
	"slices"
	"testing"
)

func TestMoveValuesByOrder(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		order    []int
		expected []int
	}{
		{
			name:     "multiple priority values",
			nums:     []int{3, 2, 5, 1, 2, 4, 1, 2},
			order:    []int{2, 1, 4},
			expected: []int{2, 2, 2, 1, 1, 4, 3, 5},
		},
		{
			name:     "priority value appears many times",
			nums:     []int{2, 3, 2, 4, 2},
			order:    []int{2},
			expected: []int{2, 2, 2, 3, 4},
		},
		{
			name:     "some priority values do not exist",
			nums:     []int{3, 2, 4},
			order:    []int{2, 7, 3},
			expected: []int{2, 3, 4},
		},
		{
			name:     "no priority values exist",
			nums:     []int{5, 6, 7},
			order:    []int{1, 2},
			expected: []int{5, 6, 7},
		},
		{
			name:     "every value is a priority value",
			nums:     []int{3, 1, 2, 1, 3},
			order:    []int{1, 3, 2},
			expected: []int{1, 1, 3, 3, 2},
		},
		{
			name:     "negative priority values",
			nums:     []int{3, -1, 2, -2, -1},
			order:    []int{-2, -1},
			expected: []int{-2, -1, -1, 3, 2},
		},
		{
			name:     "empty order",
			nums:     []int{3, 1, 4},
			order:    []int{},
			expected: []int{3, 1, 4},
		},
		{
			name:     "duplicate values in order are ignored",
			nums:     []int{3, 2, 5, 2},
			order:    []int{2, 2},
			expected: []int{2, 2, 3, 5},
		},
		{
			name:     "remaining values preserve relative order",
			nums:     []int{9, 2, 7, 1, 8, 2, 6},
			order:    []int{2},
			expected: []int{2, 2, 9, 7, 1, 8, 6},
		},
		{
			name:     "duplicate absent priority values",
			nums:     []int{3, 4, 3},
			order:    []int{7, 7, 3},
			expected: []int{3, 3, 4},
		},

		{
			name:     "duplicate remaining values should not be repeated",
			nums:     []int{5, 2, 5, 3},
			order:    []int{2},
			expected: []int{2, 5, 5, 3},
		},
		{
			name:     "duplicate remaining values separated by another value",
			nums:     []int{5, 2, 3, 5},
			order:    []int{2},
			expected: []int{2, 5, 3, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := moveValuesByOrder(tt.nums, tt.order)

			if !slices.Equal(result, tt.expected) {
				t.Errorf(
					"moveValuesByOrder(%v, %v) = %v; expected %v",
					tt.nums,
					tt.order,
					result,
					tt.expected,
				)
			}
		})
	}
}
