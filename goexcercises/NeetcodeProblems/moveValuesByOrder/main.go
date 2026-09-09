package moveValuesByOrder

import "slices"

func moveValuesByOrder(nums []int, order []int) []int {
	var results []int
	frequency := map[int]int{}
	seen := map[int]bool{}
	for _, num := range nums {
		frequency[num]++
		seen[num] = false
	}
	for _, num := range order {
		if !seen[num] {
			temp := slices.Repeat([]int{num}, frequency[num])
			results = append(results, temp...)
			seen[num] = true
		}

	}
	for _, num := range nums {
		if !seen[num] {
			results = append(results, num)
		}
	}
	return results
}
