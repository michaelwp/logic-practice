package main

func findPeakElement(nums []int) int {
	for i, n := range nums {
		if i == len(nums) - 1 {
			return i
		}

		if n > nums[i+1] {
			return i
		}
	}

	return 0
}
