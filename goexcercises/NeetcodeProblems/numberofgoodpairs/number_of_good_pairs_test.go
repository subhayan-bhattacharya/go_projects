package numberofgoodpairs

import "testing"

func TestNumberOfGoodPairs(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "first",
			nums:     []int{1, 2, 3, 1, 1, 3},
			expected: 4,
		},
		{
			name:     "pairs spread across multiple values",
			nums:     []int{4, 2, 4, 2, 4},
			expected: 4,
		},
		{
			name:     "negative numbers can form pairs",
			nums:     []int{-1, -1, 2, -1, 2},
			expected: 4,
		},
		{
			name:     "single element has no pair",
			nums:     []int{7},
			expected: 0,
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			got := numIdenticalPairs(testCase.nums)
			if got != testCase.expected {
				t.Fatalf("expected %d , got %d", testCase.expected, got)
			}
		})
	}
}
