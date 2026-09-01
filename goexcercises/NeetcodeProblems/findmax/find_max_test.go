package findmax

import "testing"

func TestFindMax(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "example",
			nums: []int{4, 2, 7, 1, 5},
			want: 7,
		},
		{
			name: "single element",
			nums: []int{42},
			want: 42,
		},
		{
			name: "all negative numbers",
			nums: []int{-8, -3, -11, -5},
			want: -3,
		},
		{
			name: "maximum is first",
			nums: []int{9, 4, 6, 2},
			want: 9,
		},
		{
			name: "maximum is last",
			nums: []int{1, 3, 5, 8},
			want: 8,
		},
		{
			name: "maximum appears more than once",
			nums: []int{6, 6, 1, 6},
			want: 6,
		},
		{
			name: "zero is the maximum",
			nums: []int{-3, 0, -1},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findMax(tt.nums)
			if got != tt.want {
				t.Fatalf("findMax(%v) = %d; want %d", tt.nums, got, tt.want)
			}
		})
	}
}
