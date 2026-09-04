package issortednondecreasing

import "testing"

func TestIsSortedNonDecreasing(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected bool
	}{
		{
			name:     "first",
			nums:     []int{1, 2, 3, 4},
			expected: true,
		},
		{
			name:     "second",
			nums:     []int{1, 3, 2, 4},
			expected: false,
		},
		{
			name:     "third",
			nums:     []int{2, 4, 4, 7, 3},
			expected: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isSorted(tt.nums)
			if got != tt.expected {
				t.Fatalf("did not work got %v", got)
			}
		})
	}
}
