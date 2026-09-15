package qualifyingSuffixLength

import "slices"

func qualifyingSuffixLength(nums []int, minimum int) int {
	var result int
	copied := append([]int(nil), nums...)
	slices.Sort(copied)
	for _, num := range copied {
		if num >= minimum {
			result++
		}
	}
	return result
}
