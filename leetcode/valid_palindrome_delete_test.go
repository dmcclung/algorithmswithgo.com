package leetcode

import (
	"fmt"
	"testing"
)

func TestValidPalindromeDelete(t *testing.T) {
	tests := []struct {
		palindrome string
		want       bool
	}{
		{
			"aba",
			true,
		},
		{
			"abca",
			true, // Could delete 'c'
		},
		{
			"abc",
			false,
		},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("Palindrome %v with delete", tc.palindrome), func(t *testing.T) {
			got := ValidPalindromeDelete(tc.palindrome)
			if got != tc.want {
				t.Fatalf("ValidPalindromeDelete() = %v wants %v", got, tc.want)
			}
		})
	}
}
