package numberofgoodpairs

func numIdenticalPairs(nums []int) int {
	frequency := map[int]int{}
	pairs := 0
	for _, num := range nums {
		pairs = pairs + frequency[num]
		frequency[num]++
	}
	return pairs
}
