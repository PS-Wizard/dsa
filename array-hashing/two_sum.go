package arrayhashing

func twoSum(nums []int, target int) []int {
	seen := make(map[int]int)
	for i, v := range nums {
		if idx, exists := seen[target-v]; exists {
			return []int{i, idx}
		}
		seen[v] = i
	}
	return []int{}
}
