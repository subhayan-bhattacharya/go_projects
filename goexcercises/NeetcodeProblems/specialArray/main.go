package specialArray

func specialArray(nums []int) int {
	result := -1
	maximum := len(nums)
	frequency := map[int]int{}
	for _, num := range nums {
		if num > maximum {
			frequency[maximum]++
		} else {
			frequency[num]++
		}
	}
	countMoreThanSelf := make([]int, len(nums)+1)
	for i := len(nums); i >= 1; i-- {
		if i == len(nums) {
			countMoreThanSelf[i] = frequency[i]
		} else {
			countMoreThanSelf[i] = countMoreThanSelf[i+1] + frequency[i]
		}
		if countMoreThanSelf[i] == i {
			result = i
			break
		}
	}
	return result
}
