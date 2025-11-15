package arrayhashing

// Initial Guess Solution
// func IsAnagram(s, t string) bool {
// 	seen := make(map[string]any)
//
// 	for v := range strings.SplitSeq(s, "") {
// 		seen[v] = 1
// 	}
//
// 	for v := range strings.SplitSeq(t, "") {
// 		if _, exists := seen[v]; !exists {
// 			return false;
// 		}
// 	}
//
// 	return true
// }

// Turns out anagrams need to have the same length and the same number of occurances for each letter too
func IsAnagram(s, t string) bool {

	if len(s) != len(t) {
		return false
	}

	seen := make(map[rune]int)

	for _, v := range s {
		seen[v]++
	}

	for _, v := range t {
		seen[v]--
		if seen[v] < 0 {
			return false
		}
	}

	return true
}
