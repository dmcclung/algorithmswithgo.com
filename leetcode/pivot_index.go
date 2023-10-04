package leetcode

func PivotIndex(nums []int) int {
	for p := 0; p < len(nums); p++ {
		// calculate left and right sums for each pivot
		var sumLeft int
		if p == 0 {
			sumLeft = 0
		} else {
			for i := 0; i < p; i++ {
				sumLeft += nums[i]
			}
		}

		var sumRight int
		if p == len(nums)-1 {
			sumRight = 0
		} else {
			for i := p + 1; i < len(nums); i++ {
				sumRight += nums[i]
			}
		}
		// if they are equal return
		if sumLeft == sumRight {
			return p
		}
	}
	return -1
}
