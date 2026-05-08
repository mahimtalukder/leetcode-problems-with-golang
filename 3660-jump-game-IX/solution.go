package main

func maxValue(nums []int) []int {
	n := len(nums)
	toMaxValue := make([]int, n)
	toMinValue := make([]int, n)
	ans := make([]int, n)

	toMaxValue[0] = nums[0]
	for i := 1; i<n; i++{
		toMaxValue[i] = max(nums[i], toMaxValue[i-1])
	}

	toMinValue[n-1] = nums[n-1]
	for i := n-2; i>=0; i--{
		toMinValue[i] = min(nums[i], toMinValue[i+1])
	}

	ans[n-1] = toMaxValue[n-1]
	for i := n-2; i>=0; i--{
		if toMaxValue[i] > toMinValue[i+1]{
			ans[i] = ans[i+1]
		}else{
			ans[i] = toMaxValue[i]
		}
	}

	return ans
}
