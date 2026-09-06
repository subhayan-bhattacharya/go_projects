package findmissing

import (
	"slices"
	"testing"
)

func TestFindMissing(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected []int
	}{
		{
			name:     "some numbers missing with duplicate",
			nums:     []int{1, 2, 2, 4},
			expected: []int{3, 5},
		},
		{
			name:     "nothing missing",
			nums:     []int{1, 2, 3, 4, 5},
			expected: []int{},
		},
		{
			name:     "several numbers missing",
			nums:     []int{2, 2, 3, 3},
			expected: []int{1, 4, 5},
		},
		{
			name:     "only one distinct number present",
			nums:     []int{5, 5, 5},
			expected: []int{1, 2, 3, 4},
		},
		{
			name:     "all values duplicated but all present",
			nums:     []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5},
			expected: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findMissing(tt.nums)

			if !slices.Equal(result, tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}
