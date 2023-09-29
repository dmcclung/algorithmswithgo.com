package leetcode

import (
	"fmt"
	"testing"
)

func TestValidPalindrome(t *testing.T) {
	tests := []struct {
		palindrome string
		want       bool
	}{
		{
			"tacocat",
			true,
		},
		{
			"abca",
			false,
		},
		{
			"abc",
			false,
		},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("Palindrome %v", tc.palindrome), func(t *testing.T) {
			got := ValidPalindrome(tc.palindrome)
			if got != tc.want {
				t.Fatalf("ValidPalindrome() = %v wants %v", got, tc.want)
			}
		})
	}
}
