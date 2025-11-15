package arrayhashing

import (
	"reflect"
	"sort"
	"testing"
)

func TestHasDuplicate(t *testing.T) {
	tests := []struct {
		nums []int
		want bool
	}{
		{[]int{2, 7, 11, 15}, false},
		{[]int{2, 2, 11, 15}, true},
		{[]int{-1, -1, -1, -2}, true},
	}

	for _, tt := range tests {
		// got := HasDuplicate(tt.nums)
		got := HasDuplicateHashMap(tt.nums)
		if got != tt.want {
			t.Errorf("Failed Contains Duplicate: %d, got: %t, want: %t", tt.nums, got, tt.want)
		}
	}
}

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		s    string
		t    string
		want bool
	}{
		{
			"hello",
			"lehhp",
			false,
		},
		{
			"xx",
			"x",
			false,
		},
	}

	for _, tt := range tests {
		// got := HasDuplicate(tt.nums)
		got := IsAnagram(tt.s, tt.t)
		if got != tt.want {
			t.Errorf("Failed Anagram Test: s: %s,t : %s, got: %t, want: %t", tt.s, tt.t, got, tt.want)
		}
	}
}

func TestTwoSum(t *testing.T) {
	tests := []struct {
		i      []int
		target int
		want   []int
	}{
		{
			[]int{3, 1, 2, 4},
			7,
			[]int{0, 3},
		},
		{
			[]int{100, 69, 3, 4, 1},
			101,
			[]int{0, 4},
		},
	}
	for _, tt := range tests {
		got := twoSum(tt.i, tt.target)

		// Sort both slices to compare
		sort.Ints(got)
		sort.Ints(tt.want)

		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("Failed Twosum Test: nums: %v, target: %d, got: %v, want: %v", tt.i, tt.target, got, tt.want)
		}
	}
}
