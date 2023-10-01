package leetcode

import (
	"fmt"
	"testing"
)

func TestBuddyStrings(t *testing.T) {
	tests := []struct {
		s    string
		goal string
		want bool
	}{
		{
			"ab",
			"ba",
			true,
		},
		{
			"ab",
			"ab",
			false,
		},
		{
			"aa",
			"aa",
			true,
		},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("Buddy string %v and goal %v", tc.s, tc.goal), func(t *testing.T) {
			got := BuddyStrings(tc.s, tc.goal)
			if got != tc.want {
				t.Fatalf("BuddyStrings() = %v wants %v", got, tc.want)
			}
		})
	}
}
