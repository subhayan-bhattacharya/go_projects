package sumfirstincreasingrun

import "testing"

func TestMaxIncreasingRunSum(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "first",
			nums:     []int{2, 4, 7, 3, 5},
			expected: 13,
		},
		{
			name:     "second",
			nums:     []int{5, 8, 8, 10},
			expected: 18,
		},
		{
			name:     "third",
			nums:     []int{5, 6, 1, 10, 20},
			expected: 31,
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			got := maxIncreasingRunSum(testCase.nums)
			if got != testCase.expected {
				t.Fatalf("expected %d , got %d", testCase.expected, got)
			}
		})
	}
}
