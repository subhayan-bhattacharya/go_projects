package pivotIndex

func pivotIndex(nums []int) int {
	result := -1
	total := 0
	for _, n := range nums {
		total += n
	}
	current := nums[0]
	if len(nums) == 1 {
		return 0
	}
	if total-current == 0 {
		return 0
	}
	for i := 1; i <= len(nums)-1; i++ {
		remaining := total - current - nums[i]
		if remaining == current {
			result = i
			break
		}
		current += nums[i]
	}
	return result
}
