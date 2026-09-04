package issortednondecreasing

func isSorted(nums []int) bool {
	isSorted := true
	prev := nums[0]
	for _, num := range nums {
		if num < prev {
			isSorted = false
			break
		}
		prev = num
	}
	return isSorted
}
