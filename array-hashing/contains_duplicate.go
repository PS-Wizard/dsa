package arrayhashing

// import "fmt"

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
