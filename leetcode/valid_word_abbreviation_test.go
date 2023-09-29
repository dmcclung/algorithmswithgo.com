package leetcode

import (
	"fmt"
	"testing"
)

func TestValidWordAbbreviation(t *testing.T) {
	tests := []struct {
		word string
		abbr string
		want bool
	}{
		{
			"internationalization",
			"i12iz4n",
			true,
		},
		{
			"apple",
			"a2e",
			false,
		},
		{
			"a",
			"01",
			false,
		},
		{
			"word",
			"3e",
			false,
		},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("%v word and abbr %v", tc.word, tc.abbr), func(t *testing.T) {
			got := ValidWordAbbreviation(tc.word, tc.abbr)
			if got != tc.want {
				t.Fatalf("ValidWordAbbreviation() = %v; want %v", got, tc.want)
			}
		})
	}
}
