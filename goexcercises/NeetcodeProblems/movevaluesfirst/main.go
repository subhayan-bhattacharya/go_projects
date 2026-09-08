package movevaluesfirst

import "slices"

func moveValuesFirst(nums []int, first int, second int) []int {
	results := make([]int, 0, len(nums))
	frequency := map[int]int{}
	for _, num := range nums {
		if num == first || num == second {
			frequency[num]++
		}
	}
	temp := slices.Repeat([]int{first}, frequency[first])
	results = append(results, temp...)
	temp = slices.Repeat([]int{second}, frequency[second])
	results = append(results, temp...)
	for _, num := range nums {
		if num != first && num != second {
			results = append(results, num)
		}
	}
	return results
}
