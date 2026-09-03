package sumfirstincreasingrun

func maxIncreasingRunSum(nums []int) int {
	result := nums[0]
	maxResult := result
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			result = result + nums[i]
		} else {
			if result > maxResult {
				maxResult = result
			}
			result = nums[i]
		}
	}
	if result > maxResult {
		maxResult = result
	}
	return maxResult
}
