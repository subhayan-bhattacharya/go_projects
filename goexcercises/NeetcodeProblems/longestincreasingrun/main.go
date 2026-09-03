package longestincreasingrun

func longestIncreasingRun(nums []int) int {
	run := 1
	longestRun := run
	previous := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > previous {
			run++
		} else {
			if run > longestRun {
				longestRun = run
			}
			run = 1
		}
		previous = nums[i]
	}
	if run > longestRun {
		longestRun = run
	}
	return longestRun
}
