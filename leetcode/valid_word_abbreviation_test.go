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
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%v word and abbr %v", test.word, test.abbr), func(t *testing.T) {
			got := ValidWordAbbreviation(test.word, test.abbr)
			if got != test.want {
				t.Fatalf("ValidWordAbbreviation() = %v; want %v", got, test.want)
			}
		})
	}
}