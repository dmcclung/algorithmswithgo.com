package leetcode

func ArrayForm(num []int, k int) []int {
    // Sum all the numbers in num
    var sum int
    for _, n := range num {
        if sum > 0 {
            sum = sum * 10
        }
        sum += n
    }

    // Add k
    sum += k
    
    // Convert number back into array-form
    var arrayForm = []int{}
    for sum > 0 {
        rem := sum % 10
        arrayForm = append([]int{rem}, arrayForm...)
        sum = sum / 10
    }
    return arrayForm   
}

func FasterArrayForm(num []int, k int) []int {
	var res = []int{}
	for i := len(num) - 1; i >= 0 || k > 0; i-- {
		if i >= 0 {
			k += num[i]
		}
		res = append([]int{k%10}, res...)
		k = k / 10
	}

	return res
}