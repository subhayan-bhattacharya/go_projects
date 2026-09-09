package specialArray

import "testing"

func TestSpecialArray(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "two elements both qualify",
			nums:     []int{3, 5},
			expected: 2,
		},
		{
			name:     "no valid x with all zeros",
			nums:     []int{0, 0},
			expected: -1,
		},
		{
			name:     "valid x is not present in array",
			nums:     []int{4, 4, 4, 4},
			expected: 4,
		},
		{
			name:     "mixed small values",
			nums:     []int{0, 4, 3, 0, 4},
			expected: 3,
		},
		{
			name:     "single qualifying element",
			nums:     []int{100},
			expected: 1,
		},
		{
			name:     "single zero has no solution",
			nums:     []int{0},
			expected: -1,
		},
		{
			name:     "answer equals array length",
			nums:     []int{5, 5, 5},
			expected: 3,
		},
		{
			name:     "large numbers do not determine x directly",
			nums:     []int{10, 20, 30, 40},
			expected: 4,
		},
		{
			name:     "some values below answer",
			nums:     []int{1, 2, 5, 6},
			expected: 2,
		},
		{
			name:     "duplicate boundary values",
			nums:     []int{2, 2},
			expected: 2,
		},
		{
			name:     "boundary prevents solution",
			nums:     []int{1, 1, 1, 2},
			expected: -1,
		},
		{
			name:     "unsorted input",
			nums:     []int{6, 0, 7, 2, 1},
			expected: 2,
		},
		{
			name:     "many zeros with a few large values",
			nums:     []int{0, 0, 0, 8, 9},
			expected: 2,
		},
		{
			name:     "one large value among zeros",
			nums:     []int{0, 0, 10, 0},
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := specialArray(tt.nums)
			if got != tt.expected {
				t.Fatalf("array : %v expected %d got %d", tt.nums, tt.expected, got)
			}
		})
	}
}
