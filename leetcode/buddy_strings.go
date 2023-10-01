package leetcode

func BuddyStrings(s, goal string) bool {
	if len(s) != len(goal) {
		return false
	}

	if s == goal {
		seen := map[rune]bool{}
		for _, c := range s {
			if seen[c] {
				return true
			} else {
				seen[c] = true
			}
		}
		return false
	}

	var diff []int
	for i := 0; i < len(s); i++ {
		if s[i] != goal[i] {
			diff = append(diff, i)
		}
	}

	return len(diff) == 2 && s[diff[0]] == goal[diff[1]] && s[diff[1]] == goal[diff[0]]
}
