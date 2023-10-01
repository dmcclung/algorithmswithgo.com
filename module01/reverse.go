package module01

// Reverse will return the provided word in reverse
// order. Eg:
//
//	Reverse("cat") => "tac"
//	Reverse("alphabet") => "tebahpla"
func TwoPointerReverse(word string) string {
	str := []rune(word)
	i, j := 0, len(str)-1

	for i < j {
		str[i], str[j] = str[j], str[i]

		i++
		j--
	}

	return string(str)
}

func Reverse(word string) string {
	var reverseStr string
	for _, r := range word {
		reverseStr = string(r) + reverseStr
	}
	return reverseStr
}
