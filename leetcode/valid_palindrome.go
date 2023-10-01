package leetcode

// Given a string s, return true if the s can be palindrome after
// deleting at most one character from it.

func ValidPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}
