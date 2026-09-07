package frequencysort

import (
	"cmp"
	"slices"
)

func frequencySort(nums []int) []int {
	result := append([]int(nil), nums...)
	frequency := map[int]int{}
	for _, num := range nums {
		frequency[num]++
	}
	slices.SortFunc(result, func(a, b int) int {
		return cmp.Or(cmp.Compare(frequency[a], frequency[b]), cmp.Compare(b, a))
	})
	return result
}
