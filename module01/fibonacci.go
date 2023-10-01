package module01

// Fibonacci returns the nth fibonacci number.
//
// A Fibonacci number N is defined as:
//
//   Fibonacci(N) = Fibonacci(N-1) + Fibonacci(N-2)
//
// Where the following base cases are used:
//
//   Fibonacci(0) => 0
//   Fibonacci(1) => 1
//
//
// Examples:
//
//   Fibonacci(0) => 0
//   Fibonacci(1) => 1
//   Fibonacci(2) => 1
//   Fibonacci(3) => 2
//   Fibonacci(4) => 3
//   Fibonacci(5) => 5
//   Fibonacci(6) => 8
//   Fibonacci(7) => 13
//   Fibonacci(14) => 377
//
func Fibonacci(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return Fibonacci(n - 1) + Fibonacci(n - 2)
	}
}

func FibonacciIter(n int) int {
	stack := []int{}
	
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		stack = append(stack, n)
	}

	var sum int
	for len(stack) > 0 {
		top := stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]

		if top == 0 || top == 1 {
			sum += top
		} else {
			stack = append(stack, top - 1)
			stack = append(stack, top - 2)
		}
	}
	return sum
}
