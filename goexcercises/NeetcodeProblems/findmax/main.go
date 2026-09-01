package findmax

func findMax(nums []int) int {
	highest := nums[0]
	for _, num := range nums {
		if num > highest {
			highest = num
		}
	}
	return highest
}
