package leetcode

import (
	"fmt"
	"strings"
)

func AlienSorted(words []string, order string) bool {
    // Turn order into a map of indexed characters
    var m = map[rune]int{}
    for i, r := range order {
        m[r] = i
    }
    
    // Convert each word to a number
    var converted = []string{}
    for _, word := range words {
        var num string
        for _, r := range word {
            i := m[r]
            num = num + fmt.Sprint(i)
        }
        converted = append(converted, num)
    }
    
    // Iterate over each word and compare to next word
    // If word i is not less than or equal to i+1, then return false
    for i := 0; i < len(converted) - 1; i++ {
        if strings.Compare(converted[i], converted[i + 1]) == 1 {
            return false
        }
    }
    
    // Otherwise return true
    return true
}
