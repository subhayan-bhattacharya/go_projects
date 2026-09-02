package findmaxfromright

func maxSeenFromRight(nums []int) []int {
	results := make([]int, len(nums))
	maxSoFar := nums[len(nums)-1]
	end := len(nums) - 1
	for start := end; start >= 0; start-- {
		num := nums[start]
		if num > maxSoFar {
			results[start] = num
			maxSoFar = num
		} else {
			results[start] = maxSoFar
		}
	}
	return results
}
