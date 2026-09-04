package mostfrequentnumber

import "testing"

func TestMostFrequentNumber(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "first",
			nums:     []int{2, 1, 2, 3, 2},
			expected: 2,
		},
		{
			name:     "second",
			nums:     []int{4, 4, 1, 1, 1, 7},
			expected: 1,
		},
		{
			name:     "third",
			nums:     []int{-2, 5, -2, -2, 5},
			expected: -2,
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			got := mostFrequent(testCase.nums)
			if got != testCase.expected {
				t.Fatalf("expected %d , got %d", testCase.expected, got)
			}
		})
	}
}
