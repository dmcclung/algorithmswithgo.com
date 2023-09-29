package leetcode

func BuddyStrings(s, goal string) bool {
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			r := []rune(s)
			temp := r[i]
			r[i] = r[j]
			r[j] = temp
			if goal == string(r) {
				return true
			}
		}
	}
	
	return false
}