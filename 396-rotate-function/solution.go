package main

func maxRotateFunction(nums []int) int {
	var maxValue, sum, current int
	n := len(nums)

	for i, num := range nums{
		sum += num
		current += i*num
	}

	maxValue = current

	for j := 1; j< n; j++ {
		current = current + sum - n* nums[n-j]
		if current > maxValue {
			maxValue = current
		}
	}

	return maxValue
}