package arrayhashing

// import "fmt"


// Initial Naive Approach
func HasDuplicate(nums []int) bool {
	seen := make([]int, len(nums))
	for _, v := range nums {
		// fmt.Printf("v: %d, seen: %v\n", v, seen)
		for _, vs := range seen {
			if v == vs {
				return true
			}
			// fmt.Printf("vs: %d, seen: %v\n", vs, seen)
		}
		seen = append(seen, v)
	}
	return false
}

// BEST APPROACH:
func HasDuplicateHashMap(nums []int) bool {
	seen := make(map[int]int)
	for _, v := range nums {
		_, exists := seen[v]
		if exists {
			return true
		}
		seen[v] = 69 // any arbitary number to mark it as seen
	}
	return false
}
