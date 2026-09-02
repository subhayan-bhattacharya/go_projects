package findmaxfromright

import (
	"slices"
	"testing"
)

func TestFindMaxFromRight(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected []int
	}{
		{
			name:     "first",
			nums:     []int{4, 2, 7, 3},
			expected: []int{7, 7, 7, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxSeenFromRight(tt.nums)
			if !slices.Equal(got, tt.expected) {
				t.Fatalf("expected the answer to be %+v , got %+v", tt.expected, got)
			}
		})
	}
}
