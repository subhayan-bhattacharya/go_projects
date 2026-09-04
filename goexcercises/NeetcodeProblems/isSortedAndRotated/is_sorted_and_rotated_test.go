package isSortedAndRotated

import "testing"

func TestIsSortedAndRotated(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected bool
	}{
		{
			name:     "first",
			nums:     []int{3, 4, 5, 1, 2},
			expected: true,
		},
		{
			name:     "second",
			nums:     []int{1, 2, 3, 4, 5},
			expected: true,
		},
		{
			name:     "third",
			nums:     []int{2, 1, 3, 4},
			expected: false,
		},
		{
			name:     "fourth",
			nums:     []int{2, 3, 1, 2},
			expected: true,
		},
		{
			name:     "fifth",
			nums:     []int{3, 1, 2, 2},
			expected: true,
		},
		{
			name:     "all equal",
			nums:     []int{2, 2, 2, 2},
			expected: true,
		},
		{
			name:     "multiple disorder points",
			nums:     []int{3, 1, 4, 2},
			expected: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isSortedAndRotated(tt.nums)
			if got != tt.expected {
				t.Fatalf("did not get : %v", tt.expected)
			}
		})
	}
}
