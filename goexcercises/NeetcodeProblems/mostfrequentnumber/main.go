package mostfrequentnumber

func mostFrequent(nums []int) int {
	frequency := map[int]int{}
	for _, num := range nums {
		frequency[num]++
	}
	highest := 0
	highestNum := nums[0]
	for k, v := range frequency {
		if v > highest {
			highest = v
			highestNum = k
		}
	}
	return highestNum
}
