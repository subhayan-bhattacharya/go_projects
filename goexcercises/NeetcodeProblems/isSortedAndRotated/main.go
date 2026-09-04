package isSortedAndRotated

func isSortedAndRotated(nums []int) bool {
	var result bool
	breaks := 0
	prev := nums[0]
	for i := 1; i < len(nums); i++ {
		current := nums[i]
		if prev > current {
			breaks++
			if breaks > 1 {
				return false
			}
		}
		prev = nums[i]

	}
	if nums[0] < nums[len(nums)-1] {
		breaks++
	}
	if breaks == 0 || breaks == 1 {
		result = true
	}
	return result
}
