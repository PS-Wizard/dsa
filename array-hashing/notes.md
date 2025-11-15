# Contains Duplicate

==Initial Solution==
```go
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
```

- This I believe is O(n^2)
    - But yeah this isn't optimal, im assuming the optimal one is somethin with a hashmap

### Better approach 1:
To sort the array first, because say we have

```
[4,3,1,7,8,8] => [1,3,4,7,8,8] sorted!
in this case, when we loop the sorted array from index 0 to 5, we need to check if 

(sorted[current_idx + 1] == sorted[current_idx])? => if so, it contains duplicate
```
- This is O(n log n), because it takes (n log n ) to sort the array first
- Better, but it can probably get better

### Best approach (hashmap; surprise lol) 
turns out the best approach is just, throwing a hashmap at the problem 
```go
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
```
- I believe this is just O(n) now, neat.

---

# Is Anagram
==My Initial Guess==

```go
func IsAnagram(s, t string) bool {
	seen := make(map[string]any)

	for v := range strings.SplitSeq(s, "") {
		seen[v] = 1
	}

	for v := range strings.SplitSeq(t, "") {
		if _, exists := seen[v]; !exists {
			return false;
		}
	}

	return true
}
```
- I assume this is just 2 * O(n), but since we drop constants -> O(n)
    - Turns out i was right
- Also turns out that anagrams need to have the same number of occurances for both the strings, didn't know that
- Final Solution
```go
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
```
- I assume this is just `2*O(N)` so again thats just O(N)

==pro tip, start with just throwing a hashmap at the problem xD==


