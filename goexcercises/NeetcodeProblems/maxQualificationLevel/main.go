package maxQualificationLevel

import (
	"slices"
)

func maxQualificationLevel(nums []int) int {
	var result int
	copied := append([]int(nil), nums...)
	slices.Sort(copied)
	greaterThanOrEqual := map[int]int{}
	length := len(copied)
	for index, num := range copied {
		numOfElements := length - index
		if num >= numOfElements {
			result = numOfElements
			break
		}
		greaterThanOrEqual[num] = numOfElements
	}
	return result
}
