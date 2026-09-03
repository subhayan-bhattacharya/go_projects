package longestincreasingrun

import "testing"

func TestLongestIncreasingRun(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "First",
			nums:     []int{1, 3, 5, 2, 4},
			expected: 3,
		},
		{
			name:     "Second",
			nums:     []int{4, 6, 6, 7, 8},
			expected: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := longestIncreasingRun(tt.nums)
			if got != tt.expected {
				t.Fatalf("expected %d and got %d", tt.expected, got)
			}
		})
	}
}
