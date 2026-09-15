package pivotIndex

import "testing"

func TestPivotIndex(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "pivot in the middle",
			nums:     []int{1, 7, 3, 6, 5, 6},
			expected: 3,
		},
		{
			name:     "pivot at first index",
			nums:     []int{2, 1, -1},
			expected: 0,
		},
		{
			name:     "pivot at last index",
			nums:     []int{-1, 1, 2},
			expected: 2,
		},
		{
			name:     "no pivot",
			nums:     []int{1, 2, 3},
			expected: -1,
		},
		{
			name:     "single element",
			nums:     []int{5},
			expected: 0,
		},
		{
			name:     "multiple possible pivots return leftmost",
			nums:     []int{0, 0, 0},
			expected: 0,
		},
		{
			name:     "negative and positive values",
			nums:     []int{-1, -1, 0, 1, 1},
			expected: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := pivotIndex(tt.nums)

			if result != tt.expected {
				t.Errorf(
					"pivotIndex(%v) = %d; expected %d",
					tt.nums,
					result,
					tt.expected,
				)
			}
		})
	}
}
