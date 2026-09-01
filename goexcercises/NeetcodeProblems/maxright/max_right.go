package maxright

// maxToRight returns the maximum value strictly to the right of each index.
//
// The final input index has no value to its right, so it has no corresponding
// entry in the result. Empty and single-element inputs therefore return an
// empty result.
func maxToRight(nums []int) []int {
	start := 0
	end := len(nums) - 1
	var results []int
	for i := start; i <= end; i++ {
		if i == len(nums)-1 {
			results = append(results, -1)
			continue
		}
		maxNumber := nums[i+1]
		for j := i + 1; j <= end; j++ {
			if nums[j] > maxNumber {
				maxNumber = nums[j]
			}
		}
		results = append(results, maxNumber)
	}
	return results
}
