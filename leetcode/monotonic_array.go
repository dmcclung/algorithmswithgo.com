package leetcode

func MonotonicArray(nums []int) bool {
	if len(nums) < 2 {
		return true // A sequence of 0 or 1 elements is trivially monotonic
	}

	increasing := true
	decreasing := true

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			increasing = false
		} else if nums[i] < nums[i+1] {
			decreasing = false
		}
	}

	return increasing || decreasing
}
