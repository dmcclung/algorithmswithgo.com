package leetcode

import "unicode"

func ValidWordAbbreviation(word string, abbr string) bool {
	i, j := 0, 0
	for i < len(word) && j < len(abbr) {
		if abbr[j] >= 'a' && abbr[j] <= 'z' {
			if word[i] != abbr[j] {
				return false
			}
			j++
			i++
		} else {
			if abbr[j] == '0' {
				return false
			}

			num := 0
			for j < len(abbr) && unicode.IsDigit(rune(abbr[j])) {
				num = num*10 + int(abbr[j]-'0')
				j++
			}
			i += num
		}
	}
	return i == len(word) && j == len(abbr)
}
