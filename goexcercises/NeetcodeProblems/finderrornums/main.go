package finderrornums

func findErrorNums(nums []int) []int {
	result := make([]int, 2)
	seen := map[int]int{}
	for i := 1; i <= len(nums); i++ {
		seen[i] = 0
	}
	for _, num := range nums {
		seen[num]++
	}
	for k, v := range seen {
		if v > 1 {
			result[0] = k
		}
		if v == 0 {
			result[1] = k
		}
	}
	return result
}
