package leetcode

import (
	"fmt"
	"testing"
)

func TestPivotIndex(t *testing.T) {
	tests := []struct {
		nums  []int
		wants int
	}{
		{
			nums:  []int{1, 7, 3, 6, 5, 6},
			wants: 3,
		},
		{
			nums:  []int{1, 2, 3},
			wants: -1,
		},
		{
			nums:  []int{2, 1, -1},
			wants: 0,
		},
		{
			nums: []int{1,0},
			wants: 0,
		},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("PivotIndex %v", tc.nums), func(t *testing.T) {
			got := PivotIndex(tc.nums)
			if got != tc.wants {
				t.Fatalf("PivotIndex() = %v wants %v", got, tc.wants)
			}
		})
	}
}
