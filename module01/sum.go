package module01

// Sum will sum up all of the numbers passed
// in and return the result
func Sum(numbers []int) int {
	var sum int

	for _, i := range numbers {
		sum += i
	}

	return sum
}
