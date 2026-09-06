package findmissing

func findMissing(nums []int) []int {
	var results []int
	seen := map[int]bool{}
	for _, num := range nums {
		seen[num] = true
	}
	for i := 1; i <= 5; i++ {
		if _, ok := seen[i]; !ok {
			results = append(results, i)
		}
	}
	return results
}
