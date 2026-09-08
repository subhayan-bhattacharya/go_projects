package movevaluesfirst

import "slices"

func moveValueFirst(nums []int, target int) []int {
	results := make([]int, 0, len(nums))
	frequency := map[int]int{}
	for _, num := range nums {
		frequency[num]++
	}
	temp := slices.Repeat([]int{target}, frequency[target])
	results = append(results, temp...)
	for _, num := range nums {
		if num != target {
			results = append(results, num)
		}
	}
	return results
}
