package maxright

import (
	"slices"
	"testing"
)

func TestMaxToRight(t *testing.T) {
	maxInt := int(^uint(0) >> 1)
	minInt := -maxInt - 1

	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "example",
			nums: []int{4, 2, 7},
			want: []int{7, 7, -1},
		},
		{
			name: "empty input",
			nums: []int{},
			want: []int{},
		},
		{
			name: "single element has nothing to its right",
			nums: []int{42},
			want: []int{-1},
		},
		{
			name: "strictly increasing",
			nums: []int{1, 2, 3, 4},
			want: []int{4, 4, 4, -1},
		},
		{
			name: "strictly decreasing",
			nums: []int{9, 7, 5, 3},
			want: []int{7, 5, 3, -1},
		},
		{
			name: "all negative values",
			nums: []int{-4, -2, -7, -9},
			want: []int{-2, -7, -9, -1},
		},
		{
			name: "duplicate maximum values",
			nums: []int{5, 5, 1, 5},
			want: []int{5, 5, 5, -1},
		},
		{
			name: "zero among negative values",
			nums: []int{-1, 0, -2, 0},
			want: []int{0, 0, 0, -1},
		},
		{
			name: "integer limits",
			nums: []int{maxInt, minInt, 0},
			want: []int{0, 0, -1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxToRight(tt.nums)
			if !slices.Equal(got, tt.want) {
				t.Fatalf("maxToRight(%v) = %v; want %v", tt.nums, got, tt.want)
			}
		})
	}
}

func TestMaxToRightDoesNotModifyInput(t *testing.T) {
	nums := []int{4, 2, 7, 1}
	original := slices.Clone(nums)

	maxToRight(nums)

	if !slices.Equal(nums, original) {
		t.Fatalf("maxToRight modified its input: got %v; want %v", nums, original)
	}
}
