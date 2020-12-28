package twoSum

func twoSum(nums []int, target int) []int {
	index := make(map[int]int, len(nums))
	for i, n := range nums {
		if j, ok := index[target-n]; ok {
			return []int{j, i}
		}
		index[n] = i
	}
	return nil
}
