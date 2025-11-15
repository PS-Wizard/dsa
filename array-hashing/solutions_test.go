package arrayhashing

import "testing"

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
		got := HasDuplicate(tt.nums)
		if got != tt.want {
			t.Errorf("Failed Contains Duplicate: %d, got: %t, want: %t", tt.nums, got, tt.want)
		}
	}
}
