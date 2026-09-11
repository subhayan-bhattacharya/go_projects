package countsAtLeast

func countsAtLeast(nums []int) []int {
	result := make([]int, len(nums)+1)
	frequency := map[int]int{}
	for _, num := range nums {
		frequency[num]++
	}
	for i := len(nums); i >= 1; i-- {
		if i == len(nums) {
			result[i] = frequency[i]
		} else {
			result[i] = result[i+1] + frequency[i]
		}
	}
	return result
}
